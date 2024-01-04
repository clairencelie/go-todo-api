package integration

import (
	"context"
	"go_todo_api/database"
	"go_todo_api/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitiateUserRepository(t *testing.T) {
	userRepository := repository.NewUserRepository()

	assert.NotNil(t, userRepository)
}

func TestGetUserById(t *testing.T) {
	db, _ := database.NewDB("./../../")

	userRepository := repository.NewUserRepository()

	ctx := context.Background()

	user, err := userRepository.Get(ctx, db, 1)

	assert.Nil(t, err)

	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "budi", user.Username)
	assert.Equal(t, "rahasia", user.Password)
	assert.Equal(t, "budi@example.xyz", user.Email)
	assert.Equal(t, "Budi Gemilang", user.Name)
}
