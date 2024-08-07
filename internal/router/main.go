package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"svc/proxy-service/internal/middleware"
	"svc/proxy-service/internal/rest"
)

var router = gin.Default()

func Init(port string) error {
	addMiddleware()
	addRoutes()
	return router.Run(":" + port)
}

func addRoutes() {
	// direct access to database
	dbRoute := router.Group("/db",
		middleware.TokenParser(),
	)
	addDbRoutes(dbRoute)

	// proxy handle
	router.Any("/api/*path",
		middleware.TokenParser(),
		rest.ProxyHandler(),
	)

	// health check
	router.GET("/ping",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		},
	)
}

func addMiddleware() {
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorRecovery())
}
