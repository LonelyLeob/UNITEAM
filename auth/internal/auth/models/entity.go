package models

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mileusna/useragent"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	debug = os.Getenv("DEBUG")
)

type User struct {
	Id    uuid.UUID   `json:"uuid"`
	Name  string      `json:"name"`
	Pwd   string      `json:"-"`
	Email string      `json:"email"`
	Meta  []*UserMeta `json:"meta,omitempty"`
}

func SetUser(cred ...string) *User {
	return &User{
		Id:    uuid.New(),
		Name:  cred[0],
		Pwd:   cred[1],
		Email: cred[2],
		Meta:  []*UserMeta{},
	}
}

func (u *User) InputValidation() error {
	if u.Name == "" || u.Email == "" || u.Pwd == "" {
		return errValidation
	}

	return nil
}

func (u *User) HashPassword(pwd string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errWrongPassword
	}
	u.Pwd = string(hash)

	return nil
}

func (u *User) ComparePassword(pwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Pwd), []byte(pwd)); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errWrongPassword
	}

	return nil
}

// meta is entrypoint in service of concrete user
type UserMeta struct {
	Id      uuid.UUID `json:"-"`
	Lv      int64     `json:"lv"`
	Browser string    `json:"bwr"`
	OS      string    `json:"os"`
	Refresh string    `json:"-"`
}

func (um *UserMeta) ParseUserAgent(ua string) error {
	if ua == "" {
		return errNoUserAgent
	}

	parseua := useragent.Parse(ua)
	um.Browser = parseua.Name
	um.OS = parseua.OS

	return nil
}

type AccessClaims struct {
	*jwt.RegisteredClaims
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RefreshClaims struct {
	*jwt.RegisteredClaims
	Browser string `json:"browser"`
	OS      string `json:"os"`
}
