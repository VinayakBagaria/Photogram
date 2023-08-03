package restutil

import (
	"github.com/gin-gonic/gin"
)

func WriteAsJson(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, data)
}

func WriteError(c *gin.Context, statusCode int, err error) {
	e := "error"
	if err != nil {
		e = err.Error()
	}
	WriteAsJson(c, statusCode, gin.H{"error": e})
}
