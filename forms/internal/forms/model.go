package forms

import (
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Form struct {
	Uuid        uuid.UUID
	Name        string
	Description string
	IsAnonym    bool
	AuthorName  string
	Fields      []*Field
}

type Field struct {
	FormUuid *uuid.UUID
	Id       uint64
	Quiz     string
	Answers  []*VariableAnswers
}

type VariableAnswers struct {
	Id      uint64
	FieldId uint64
	Answer  string
}

func (a *VariableAnswers) ParseFieldId(sid string) error {
	id, err := strconv.ParseUint(sid, 10, 32)
	if err != nil {
		return err
	}

	a.FieldId = id

	return nil
}

type AuthClaims struct {
	jwt.RegisteredClaims
	Name string `json:"name"`
}
