package handlers

import (
	"github.com/dickeyy/passwords/api/middleware"
	"github.com/dickeyy/passwords/api/storage"
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

	if userID == nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	user, err := storage.GetUser(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	if user == nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"user": user,
		},
	})
}
