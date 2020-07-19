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
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Find(t *testing.T) {
	db, tearDown := sqlstore.TestDB(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
