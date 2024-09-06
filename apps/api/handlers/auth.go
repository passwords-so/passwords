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
	r := router.Group("/auth")
	{
		r.POST("/login", Login)
		r.POST("/register", Register)
	}
}

// POST /api/auth/login
type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	// get the json body
	var user LoginBody
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

	// check if the user exists
	if taken, _ := storage.IsEmailTaken(c, user.Email); !taken {
		c.JSON(400, gin.H{
			"message": "email not found",
		})
		return
	}

	// get the user by email
	userInDB, err := storage.GetUserByEmail(c, user.Email)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user by email")
		c.JSON(500, gin.H{
			"message": "failed to get user",
		})
		return
	}

	// validate the password
	valid, err := lib.ComparePasswordAndHash(user.Password, userInDB.Password)
	if err != nil {
		log.Error().Err(err).Msg("failed to compare passwords")
		c.JSON(500, gin.H{
			"message": "failed to compare passwords",
		})
		return
	}

	if !valid {
		c.JSON(400, gin.H{
			"message": "invalid password",
		})
		return
	}

	// generate the jwt
	jwt, err := lib.GenerateJWT(userInDB.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate jwt")
		c.JSON(500, gin.H{
			"message": "failed to generate jwt",
		})
		return
	}

	// return the user
	c.JSON(200, gin.H{
		"user":    gin.H{"id": userInDB.ID, "email": userInDB.Email},
		"token":   jwt,
		"message": "login successful",
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

	// generate the jwt
	jwt, err := lib.GenerateJWT(userModel.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate jwt")
		c.JSON(500, gin.H{
			"message": "failed to generate jwt",
		})
		return
	}

	// return the user
	c.JSON(200, gin.H{
		"user":    gin.H{"id": userID, "email": user.Email},
		"token":   jwt,
		"message": "user created",
	})
}
