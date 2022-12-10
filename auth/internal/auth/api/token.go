package api

import (
	"authenticate/internal/auth/models"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type TokenGiver struct {
	AccessDuration  time.Duration
	RefreshDuration time.Duration
	SigningKey      []byte
}

type TokenSerializer struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func NewGiver(adur, rdur time.Duration, skey string) *TokenGiver {
	return &TokenGiver{
		AccessDuration:  adur,
		RefreshDuration: rdur,
		SigningKey:      []byte(skey),
	}
}

// create map with access and refresh tokens
func (t *TokenGiver) CreatePairToken(user *models.User) (*TokenSerializer, error) {

	// access token
	iata := time.Now()
	expa := time.Now().Add(time.Duration(t.AccessDuration))

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.AccessClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expa),
			IssuedAt:  jwt.NewNumericDate(iata),
			Issuer:    "UNIAUTH",
		},
		Name:  user.Name,
		Email: user.Email,
	})

	access, err := atoken.SignedString(t.SigningKey)
	if err != nil {
		if os.Getenv("DEBUG") == "True" {
			logrus.Error(err)
		}

		return nil, ErrCantSignString
	}

	// refresh token
	iatr := time.Now()
	expr := time.Now().Add(time.Duration(t.RefreshDuration))
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.RefreshClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expr),
			IssuedAt:  jwt.NewNumericDate(iatr),
			Issuer:    "UNIAUTH",
		},
		OS:      user.Meta[0].OS,
		Browser: user.Meta[0].Browser,
	})

	refresh, err := rtoken.SignedString(t.SigningKey)
	if err != nil {
		if os.Getenv("DEBUG") == "True" {
			logrus.Error(err)
		}
		return nil, ErrCantSignString
	}

	return &TokenSerializer{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func (t *TokenGiver) RefreshAccess(refresh string) {}

func (t *TokenGiver) ParseAccess(header string) (string, error) {
	tokenstr, err := t.ParseHeader(header)
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseWithClaims(tokenstr, &models.AccessClaims{}, func(tj *jwt.Token) (interface{}, error) {
		if _, ok := tj.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedMethod
		}

		return t.SigningKey, nil
	})
	if err != nil {
		return "", ErrParseConflict
	}

	claims, ok := token.Claims.(*models.AccessClaims)
	if ok && token.Valid {
		return claims.Name, nil
	}

	return "", ErrInfo
}

func (t *TokenGiver) ParseHeader(header string) (string, error) {
	headerparts := strings.Split(header, " ")
	if len(headerparts) > 2 || headerparts[0] != "Bearer" {
		return "", ErrHeaderInvalid
	}
	return headerparts[1], nil
}
