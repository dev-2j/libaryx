package connx

import (
	"fmt"
	"net/http"

	"example.com/m/service/errorx"
	"github.com/gin-gonic/gin"
)

type GX struct {
	GinEngine      *gin.Engine
	GinRouterGroup *gin.RouterGroup
}

type GM struct {
	domainName  string
	moduleName  string
	serviceName string
	port        int
}

func New(domainName, moduleName, serviceName string, port int) *GM {
	return &GM{
		domainName:  domainName,
		moduleName:  moduleName,
		serviceName: serviceName,
		port:        port,
	}
}

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
