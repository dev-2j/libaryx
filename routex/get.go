package routex

import "github.com/gin-gonic/gin"

func GET(g *gin.RouterGroup, relativePath string, handlers ...gin.HandlerFunc) {
	path := g.BasePath()
	if relativePath != "" && relativePath != "/" {
		path = path + "/" + relativePath
	}
	g.GET(relativePath, append(handlers, Guard(path))...)
}
