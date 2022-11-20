package internal

import (
	"errors"
	"strings"
)

func GetAuthString(header string) (string, error) {
	parts := strings.Split(header, " ")
	if len(parts) == 2 {
		if parts[0] == "Bearer" {
			return parts[1], nil
		}
	}

	return "", errors.New("invalid or nothing in header, please relog to be authorize user")
}
