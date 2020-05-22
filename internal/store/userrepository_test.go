package store_test

import (
	"testing"

	"github.com/M1crogravity/go-resapi/internal/app/model"
	"github.com/M1crogravity/go-resapi/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, tearDown := store.TestStore(t, databaseURL)
	defer tearDown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, tearDown := store.TestStore(t, databaseURL)
	defer tearDown("users")

	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@example.com",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
