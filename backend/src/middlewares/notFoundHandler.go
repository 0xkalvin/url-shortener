package middlewares

import (
	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Page not found"})
}