package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mtfelian/test_task_3/storage"
)

// Service represents the service internal components
type Service struct {
	Storage    storage.Keeper
	HTTPServer *gin.Engine
}

var singleton *Service

// Get provides access to a service components
func Get() *Service { return singleton }

// New creates new service object
func New() {
	var keeper storage.Keeper
	keeper = storage.NewPostgres()
	singleton = &Service{
		Storage:    keeper,
		HTTPServer: gin.Default(),
	}
}
