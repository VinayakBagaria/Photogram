package restutil

import (
	"github.com/VinayakBagaria/photogram/dto"
	"github.com/gin-gonic/gin"
)

func WriteAsJson(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, data)
}

func WriteError(c *gin.Context, statusCode int, err error, data gin.H) {
	errorObject := dto.GeneralErrorResponse{Error: err.Error(), Meta: data}
	WriteAsJson(c, statusCode, errorObject)
}
