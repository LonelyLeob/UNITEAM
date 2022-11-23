package internal

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func createAccessToken(user *User, skey interface{}) (string, error) {
	iat := time.Now()
	exp := time.Now().Add(time.Duration(15 * time.Minute))

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &AccessClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(iat),
			Issuer:    "UNIAUTH",
		},
		Name:  user.name,
		Email: user.email,
		Role:  user.role,
	})

	access, err := atoken.SignedString(skey)
	if err != nil {
		return "", err
	}

	return access, nil
}

func createRefreshToken(user *User, skey interface{}) (string, error) {
	iat := time.Now()
	exp := time.Now().Add(time.Duration(72 * time.Hour))

	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &AccessClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(iat),
			Issuer:    "UNIAUTH",
		},
		Name: user.name,
	})

	refresh, err := rtoken.SignedString(skey)
	if err != nil {
		return "", err
	}

	return refresh, nil

}

func toJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}

func errJSON(w http.ResponseWriter, code int, err error) {
	toJSON(w, code, map[string]string{"error": err.Error()})
}

func sendCredentials(w http.ResponseWriter, atoken, rtoken string) {
	toJSON(w, http.StatusOK, map[string]string{
		"access":  atoken,
		"refresh": rtoken,
	})
}

func ParseTokenFromHeader(tokenstr string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(tokenstr, &AccessClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnexpectedMethod
		}

		return signingKey, nil
	})
	if err != nil {
		return "", errParseConflict
	}

	claims, ok := token.Claims.(*AccessClaims)
	if ok && token.Valid {
		return claims.Name, nil
	}

	return "", errInfo
}
