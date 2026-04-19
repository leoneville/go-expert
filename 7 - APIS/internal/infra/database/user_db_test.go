package database

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/go-openapi/testify/v2/assert"
	"github.com/leoneville/goexpert/api/internal/entity"
	"gorm.io/gorm"
)

var (
	name     = "John Doe"
	email    = "tw@jd.com"
	password = "123456"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser(name, email, password)
	userDB := NewUserRepository(db)

	err = userDB.Create(user)
	assert.NoError(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser(name, email, password)
	userDB := NewUserRepository(db)

	err = userDB.Create(user)
	assert.NoError(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.NoError(t, err)

	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
