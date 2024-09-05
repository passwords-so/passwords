package router

import (
	"github.com/dickeyy/passwords/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		handlers.RegisterAuthRoutes(api)
	}

	return router
}
