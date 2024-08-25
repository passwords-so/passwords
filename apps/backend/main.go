package main

import (
	"net/http"
	"os"

	"github.com/dickeyy/passwords/backend/log"
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
		log.Info().Msgf("environment set to %s", enviornment)
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

	// create engine instance
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// run server
	log.Info().Msgf("starting server on port %s", os.Getenv("PORT"))
	app.Run(":" + os.Getenv("PORT"))
}
