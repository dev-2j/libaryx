package errorx

import (
	"github.com/gin-gonic/gin"
)

// New ฟังก์ชันที่สร้างข้อผิดพลาดใหม่
func New(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}

// ErrorResponse โครงสร้างสำหรับข้อผิดพลาด
type ErrorResponse struct {
	Message string `json:"message"`
}

func (e *ErrorResponse) SetContext(c *gin.Context) *ErrorResponse {
	// เพิ่มข้อมูลที่เกี่ยวข้องกับ context
	return e
}

func (e *ErrorResponse) StatusBadRequest() *ErrorResponse {
	// กำหนดสถานะข้อผิดพลาด
	return e
}
