package handlers

import (
	"github.com/dickeyy/passwords/api/lib"
	"github.com/dickeyy/passwords/api/log"
	"github.com/dickeyy/passwords/api/structs"

	"github.com/dickeyy/passwords/api/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", Login)
		auth.POST("/register", Register)
	}
}

// POST /api/auth/login
func Login(c *gin.Context) {
	// just return hello world for now
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

// POST /api/auth/register
type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	// get the json body
	var user RegisterBody
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	// validate all fields
	if user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{
			"message": "missing fields",
		})
		return
	}

	// check if the user already exists
	if taken, _ := storage.IsEmailTaken(c, user.Email); taken {
		c.JSON(400, gin.H{
			"message": "email already taken",
		})
		return
	}

	// model the data
	userID := uuid.New().String()
	pwHash, err := lib.HashPassword(user.Password)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		c.JSON(500, gin.H{
			"message": "failed to hash password",
		})
		return
	}

	userModel := structs.User{
		ID:       userID,
		Email:    user.Email,
		Password: pwHash,
	}

	// create the user
	err = storage.CreateUser(c, &userModel)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		c.JSON(500, gin.H{
			"message": "failed to create user",
		})
		return
	}

	// just return hello world for now
	c.JSON(200, gin.H{
		"user":    gin.H{"id": userID, "email": user.Email},
		"message": "user created",
	})
}
