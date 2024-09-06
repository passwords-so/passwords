package handlers

import (
	"github.com/dickeyy/passwords/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	r := router.Group("/user")
	{
		r.GET("/me", middleware.JWTAuth(), GetMe)
	}
}

// GET /api/user/me
func GetMe(c *gin.Context) {
	userID, _ := c.Get("user_id")

	c.JSON(200, gin.H{
		"message": "hi",
		"user_id": userID,
	})
}
