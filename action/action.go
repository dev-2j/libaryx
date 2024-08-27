package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewFunction ฟังก์ชันที่ใช้จัดการคำขอ GET
func NewFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from NewFunction"})
}

// NewFunction2 ฟังก์ชันใหม่ที่จัดการคำขอ POST
func NewFunction2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from NewFunction2"})
}
