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
	Meta     *UserMeta
}

type UserMeta struct {
	Device       string
	LastSign     int64
	ActivityRate int
	Refresh      string
}

type AccessClaims struct {
	*jwt.RegisteredClaims
	Name string `json:"name"`
	Role string `json:"role"`
}

type RefreshClaims struct {
	*jwt.RegisteredClaims
	Email  string `json:"email"`
	Device string `json:"device"`
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
