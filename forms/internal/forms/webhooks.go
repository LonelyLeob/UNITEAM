package forms

import (
	"encoding/json"
	"net/http"
)

func Auth_GetAttrs(url string) (string, int, error) {
	type RespUserDTO struct {
		Name string `json:"user"`
	}
	res, err := http.Get(url)
	if err != nil {
		return "", http.StatusConflict, err
	}

	urep := &RespUserDTO{}
	if err := json.NewDecoder(res.Body).Decode(urep); err != nil {
		return "", http.StatusConflict, err
	}

	return urep.Name, http.StatusOK, nil
}
