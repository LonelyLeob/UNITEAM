package forms

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

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

func ParseTokenFromHeader(header string, signingKey []byte) (string, error) {
	token := header
	if token == "" {
		return "", errAuthHeaderNotFound
	}

	headerparts := strings.Split(token, " ")
	if headerparts[0] != "Bearer" || len(headerparts) > 2 {
		return "", errAuthHeaderInvalid
	}
	tokenstr := headerparts[0]
	auth, err := jwt.ParseWithClaims(tokenstr, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnexpectedMethod
		}

		return signingKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := auth.Claims.(*AuthClaims)
	if ok && auth.Valid {
		return claims.Username, nil
	}

	return "", errors.New("cant show info")
}
