package action

import (
	"github.com/gin-gonic/gin"
)

func XDevTest(c *gin.Context) (any, error) {

	datax := "test"
	return datax, nil
}
