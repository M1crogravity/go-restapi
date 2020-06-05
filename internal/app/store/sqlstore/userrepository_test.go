package sqlstore_test

import (
	"testing"

	"github.com/M1crogravity/go-resapi/internal/app/model"
	"github.com/M1crogravity/go-resapi/internal/app/store"
	"github.com/M1crogravity/go-resapi/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, tearDown := sqlstore.TestDB(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, tearDown := sqlstore.TestDB(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
