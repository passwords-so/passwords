package storage

import (
	"context"

	"github.com/dickeyy/passwords/api/log"
	"github.com/dickeyy/passwords/api/services"
	"github.com/dickeyy/passwords/api/structs"
	"github.com/google/uuid"
)

// CRUD
func CreateUser(ctx context.Context, user *structs.User) error {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	log.Debug().Msgf("creating user with id %s", user.ID)
	_, err := services.DB.Model(user).Insert()
	return err
}

func GetUser(ctx context.Context, id int) (*structs.User, error) {
	var user structs.User
	err := services.DB.Model(&structs.User{}).Where("id = ?", id).Select(&user)
	return &user, err
}

func GetUserByEmail(ctx context.Context, email string) (*structs.User, error) {
	var user structs.User
	err := services.DB.Model(&structs.User{}).Where("email = ?", email).Select(&user)
	return &user, err
}

func GetUsers(ctx context.Context) ([]*structs.User, error) {
	var users []*structs.User
	err := services.DB.Model(&structs.User{}).Select(&users)
	return users, err
}

func UpdateUser(ctx context.Context, user *structs.User) error {
	_, err := services.DB.Model(&structs.User{}).Where("id = ?", user.ID).Update(user)
	return err
}

// Helpers
func IsEmailTaken(ctx context.Context, email string) (bool, error) {
	var user structs.User
	err := services.DB.Model(&structs.User{}).Where("email = ?", email).Select(&user)
	if err != nil {
		return false, err
	}
	return user.ID != "", nil
}
