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
	accessDuration  time.Duration
	refreshDuration time.Duration
	signingKey      []byte
}

type TokenSerializer struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func NewGiver(adur, rdur time.Duration, skey string) *TokenGiver {
	return &TokenGiver{
		accessDuration:  adur,
		refreshDuration: rdur,
		signingKey:      []byte(skey),
	}
}

// create map with access and refresh tokens
func (t *TokenGiver) CreatePairToken(user *models.User) (*TokenSerializer, error) {

	// access token
	iata := time.Now()
	expa := time.Now().Add(time.Duration(t.accessDuration))

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.AccessClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expa),
			IssuedAt:  jwt.NewNumericDate(iata),
			Issuer:    "UNIAUTH",
		},
		Name:  user.Name,
		Email: user.Email,
	})

	access, err := atoken.SignedString(t.signingKey)
	if err != nil {
		if os.Getenv("DEBUG") == "True" {
			logrus.Error(err)
		}

		return nil, errCantSignString
	}

	// refresh token
	iatr := time.Now()
	expr := time.Now().Add(time.Duration(t.refreshDuration))
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.RefreshClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expr),
			IssuedAt:  jwt.NewNumericDate(iatr),
			Issuer:    "UNIAUTH",
		},
		OS:      user.Meta[0].OS,
		Browser: user.Meta[0].Browser,
	})

	refresh, err := rtoken.SignedString(t.signingKey)
	if err != nil {
		if os.Getenv("DEBUG") == "True" {
			logrus.Error(err)
		}
		return nil, errCantSignString
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
			return nil, errUnexpectedMethod
		}

		return t.signingKey, nil
	})
	if err != nil {
		return "", errParseConflict
	}

	claims, ok := token.Claims.(*models.AccessClaims)
	if ok && token.Valid {
		return claims.Name, nil
	}

	return "", errInfo
}

func (t *TokenGiver) ParseHeader(header string) (string, error) {
	headerparts := strings.Split(header, " ")
	if len(headerparts) > 2 || headerparts[0] != "Bearer" {
		return "", errHeaderInvalid
	}
	return headerparts[1], nil
}
