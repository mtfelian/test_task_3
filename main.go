package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mtfelian/test_task_3/api"
	"github.com/mtfelian/test_task_3/config"
	"github.com/mtfelian/test_task_3/service"
	"github.com/spf13/viper"
)

func main() {
	// initialize configuration
	if err := config.Parse(); err != nil {
		log.Fatalln(err)
	}

	// initialize service
	service.New()
	s := service.Get()

	// register and run a HTTP server
	RegisterHTTPAPIHandlers(s.HTTPServer)
	if err := s.HTTPServer.Run(fmt.Sprintf(":%d", viper.GetInt(config.Port))); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}

// RegisterHTTPAPIHandlers registers HTTP API handlers
func RegisterHTTPAPIHandlers(router *gin.Engine) {
	router.POST("/param", api.GetParam)
}
