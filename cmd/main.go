package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"svc/proxy-service/internal/config"
	"svc/proxy-service/internal/data"
	"svc/proxy-service/internal/router"
)

func main() {
	// Set up
	setup()

	// Server
	port := config.GetConfig().Server.HttpPort
	err := router.Init(port)
	if err != nil {
		return
	}
}

func setup() {
	gin.ForceConsoleColor()
	viper.AutomaticEnv()
	config.Init(viper.GetString("ENV"))
	data.ConnectDatabase()
}
