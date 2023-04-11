package middleware

import (
	"example1/app/model/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth != "castles" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  responses.Error,
			"message": "No access.",
		})
		c.Abort()
		return
	}
	c.Next()
}
