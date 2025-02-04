package user

import (
	"context"
	"log"
	"pockethealth/internchallenge/pkg/datastore"
)

// ApiServices are an instance of an ApiServicer that implement the api actions for an ApiServicer
type UserApiService struct{}

// NewUserApiService creates a new Api Service
func NewUserApiService() UserApiService {
	return UserApiService{}
}

// PostRegister - Register a User
func (s UserApiService) PostRegister(ctx context.Context, name string, email string, favouriteColor string) (string, error) {
	// save user to datastore
	userId, err := datastore.CreateUser(ctx, name, email, favouriteColor)
	if err != nil {
		log.Printf("error creating user: %q", err.Error())
		return "", err
	}

	log.Printf("created user with name: %s\n", name)
	log.Printf("created user with id: %s\n", userId)
	log.Printf("user favourite color: %s\n", favouriteColor)

	// return the user id
	// Fix the bug where it returns a empty string
	return userId, err
}
