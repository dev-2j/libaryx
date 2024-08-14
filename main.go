package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Global middleware
	r.Use(func(c *gin.Context) {
		// Do some processing here
		c.Next()
	})

	// Example of middleware that runs before the route handler
	// r.GET("/ping", action.XDevTest)

	// Grouping routes with middleware
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user": "password",
	}))
	authorized.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Admin Page")
	})

	r.Run(":8080")
}
