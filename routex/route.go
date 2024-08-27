package routex

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"example.com/m/action"
	"example.com/m/service/config"
	"example.com/m/service/errorx"
	"example.com/m/service/logx"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken(r *http.Request) (*jwt.Token, error) {

	// tokenString := extractToken(r)
	// if tokenString == "" {
	// 	logx.Alert("Token is not valid")
	// }

	var tokenString string
	bearToken := r.Header.Get("Authorization")

	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		tokenString = strArr[1]
	}
	if tokenString == `` {
		return nil, fmt.Errorf("invalid Authorization=%s", bearToken)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("VerifyToken: Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	// logx.Infoln("VerifyToken: Pass.")
	return token, nil
}

func TokenValid(r *http.Request) (*jwt.Token, error) {
	token, err := VerifyToken(r)
	if err != nil {
		// logx.Alert("VerifyToken: failure, %s", err.Error())
		return nil, err
	}
	// if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
	// 	return nil, err
	// }
	if !token.Valid {
		return nil, fmt.Errorf(`invalid token`)
	}

	// logx.Infoln("TokenValid: Pass.")
	return token, nil
}

// Guard เป็น middleware ฟังก์ชันที่ใช้ในเส้นทาง
func Guard() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ตรวจสอบ JWT token ใน Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		// ตรวจสอบความถูกต้องของ JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return []byte(config.GetEnv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			logx.Alert(`IP:%s [Invalid Token], error: %v`, c.ClientIP(), err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// ดึงข้อมูลจาก claims ใน token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// ตรวจสอบข้อมูลการเข้าสู่ระบบจาก claims
		username, ok := claims["username"].(string)
		if !ok || username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
			c.Abort()
			return
		}

		// เพิ่มข้อมูลลง context สำหรับใช้งานใน handler ต่อไป
		c.Set("username", username)

		// ตรวจสอบสิทธิ์และเรียก c.Next() เพื่อให้ handler ทำงานต่อ
		c.Next()
	}
}

// GET ใช้ในการเพิ่มเส้นทาง GET ให้กับ gin.RouterGroup
func GET(g *gin.RouterGroup, relativePath string, handlers ...gin.HandlerFunc) {
	path := g.BasePath()
	if relativePath != "" && relativePath != "/" {
		path = path + "/" + relativePath
	}
	g.GET(relativePath, handlers...)
}

type GX struct {
	serviceName    string
	GinEngine      *gin.Engine
	GinRouterGroup *gin.RouterGroup
}

type GM struct {
	domainName  string
	moduleName  string
	serviceName string
	port        int
}

// New ใช้ในการสร้าง GM ใหม่
func New(domainName, moduleName, serviceName string, port int) *GM {
	return &GM{
		domainName:  domainName,
		moduleName:  moduleName,
		serviceName: serviceName,
		port:        port,
	}
}

// Group ใช้ในการสร้างกลุ่มเส้นทาง
func (s *GX) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return s.GinRouterGroup.Group(relativePath, handlers...)
}

// Connect ใช้ในการเริ่มต้นเซิร์ฟเวอร์และการกำหนดเส้นทาง
func (s *GM) Connect(routerx func(*GX)) {
	gx := GX{}
	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	gx.GinEngine = router
	gx.GinRouterGroup = router.Group(s.moduleName)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome"})
	})

	routerx(&gx)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, errorx.New("Route not found").SetContext(c).StatusBadRequest())
	})

	_ = router.Run(fmt.Sprintf(":%v", s.port))
}

// Route ใช้ในการกำหนดเส้นทางทั้งหมดสำหรับโมดูลนี้
func Route() {
	var GO_DEV = New("example.com", "module", "service", 4000)

	GO_DEV.Connect(func(g *GX) {

		rGroup := g.Group(``)
		{
			// สร้างเส้นทาง GET ที่ /action/new
			GET(rGroup, `action/new`, action.NewFunction)
		}

		// สร้างเส้นทางที่มีการตรวจสอบสิทธิ์ด้วย Guard middleware
		rAuthen := g.Group(``, Guard())
		{
			GET(rAuthen, `action/new2`, action.NewFunction2)
		}
	})
}
