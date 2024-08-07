package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"svc/proxy-service/internal/common"
)

func ErrorRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %s\n", err)

				// Client error
				if clientError, ok := err.(common.ClientError); ok {
					c.JSON(clientError.Code, gin.H{
						"message": clientError.Message,
					})
					c.Abort()
					return
				}

				// Server error
				if serverError, ok := err.(common.ServerError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": serverError.Message,
						"detail":  serverError.Detail,
					})
					c.Abort()
					return
				}

				// Default error
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
				})
				c.Abort()
			}
		}()
		// Continue down the chain to the next handler
		c.Next()
	}
}
