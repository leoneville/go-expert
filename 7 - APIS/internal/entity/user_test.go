package entity

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

var (
	name     = "John Doe"
	email    = "j@j.com"
	password = "123456"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(name, email, password)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser(name, email, password)

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(password))
	assert.False(t, user.ValidatePassword("secretpwd"))
	assert.NotEqual(t, user.Password, password)
}
