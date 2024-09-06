package main

import (
	"context"
	"os"

	"github.com/dickeyy/passwords/api/log"
	"github.com/dickeyy/passwords/api/router"
	"github.com/dickeyy/passwords/api/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	enviornment = "dev"
)

func init() {
	log.Info().Msg("starting up")

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Error().Err(err).Msg("failed to load .env file")
	}

	// set environment
	if os.Getenv("ENVIRONMENT") != "" {
		enviornment = os.Getenv("ENVIRONMENT")
		log.Debug().Msgf("environment set to %s", enviornment)
	}

}

func main() {

	// set gin mode based on environment
	switch enviornment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// connect services
	services.ConnectDB(context.Background())
	defer services.CloseDB()

	// config router instance
	r := router.SetupRouter()

	// run server
	log.Info().Msgf("server started on port %s", os.Getenv("PORT"))
	r.Run(":" + os.Getenv("PORT"))
}
