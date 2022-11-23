package internal

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	name     string
	password string
	email    string
	role     string
}

type AccessClaims struct {
	*jwt.RegisteredClaims
	Name string `json:"name"`
	Role string `json:"role,omitempty"`
}

type RefreshClaims struct {
	*jwt.RegisteredClaims
	Email string `json:"email"`
}

func (u *User) HashPassword(password string) error {
	toByte := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(toByte, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.password = string(hash)

	return nil
}

func (u *User) CheckForRequiredParams() error {
	if u.name == "" || u.password == "" || u.email == "" {
		return errFieldsMustBeNotEmpty
	}
	return nil
}
