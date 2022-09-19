package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONRequest[T any] struct {
	Data T
}

func handleError(c *gin.Context, statusCode int, err error) {
	c.AbortWithStatusJSON(
		statusCode,
		gin.H{
			"success": false,
			"status":  statusCode,
			"error":   err.Error(),
		},
	)
}

func handleOK(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"status":  http.StatusOK,
			"data":    data,
		},
	)
}
