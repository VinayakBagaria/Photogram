package restutil

import (
	"github.com/gin-gonic/gin"
)

func WriteAsJson(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, data)
}

func WriteError(c *gin.Context, statusCode int, err error, data gin.H) {
	if data == nil {
		data = gin.H{}
	}
	data["error"] = err.Error()
	WriteAsJson(c, statusCode, data)
}
