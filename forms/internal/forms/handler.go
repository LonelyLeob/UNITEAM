package forms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (s *Server) FormsCreatingHttp() http.HandlerFunc {
	type ReqFormsDTO struct {
		Name   string `json:"name"`
		Desc   string `json:"desc"`
		Anonym bool   `json:"anon"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		headerparts := strings.Split(r.Header.Get("Authorization"), " ")

		name, code, err := Auth_GetAttrs(fmt.Sprintf("http://authenticate:7000/api/v1/attrs?token=%s", headerparts[1]))
		if err != nil {
			errJSON(w, code, err)
		}

		freq := &ReqFormsDTO{}
		if err := json.NewDecoder(r.Body).Decode(freq); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		if code == http.StatusOK {
			fuid := uuid.New()
			f := &Form{
				Uuid:        fuid,
				Name:        freq.Name,
				Description: freq.Desc,
				IsAnonym:    freq.Anonym,
				AuthorName:  name,
			}

			if err := s.store.Forms().CreateForm(f); err != nil {
				errJSON(w, http.StatusUnprocessableEntity, err)
				return
			} else {
				toJSON(w, http.StatusCreated, f)
			}
		} else {
			errJSON(w, code, errCodeIsNotOK)
			return
		}
	}
}

func (s *Server) FieldCreatingForm() http.HandlerFunc {
	type request struct {
		QuizName string `json:"quiz"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := ParseTokenFromHeader(r.Header.Get("Authorization"), s.signingKey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, errBadToParseToken)
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		form := r.URL.Query().Get("form")

		fiuid, err := uuid.Parse(form)
		if err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		f := &Field{
			FormUuid: &fiuid,
			Quiz:     req.QuizName,
		}

		if err := s.store.Forms().AddField(f); err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		} else {
			toJSON(w, http.StatusCreated, f)
		}
	}
}

func (s *Server) AnswerCreatingField() http.HandlerFunc {
	type Request struct {
		Answer string `json:"answer"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := ParseTokenFromHeader(r.Header.Get("Authorization"), s.signingKey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, errBadToParseToken)
			return
		}

		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		field := r.URL.Query().Get("field")

		a := &VariableAnswers{
			Answer: req.Answer,
		}

		if err := a.ParseFieldId(field); err != nil {
			errJSON(w, http.StatusBadRequest, errParseInt)
			return
		}

		if err := s.store.Forms().AddAnswer(a); err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		} else {
			toJSON(w, http.StatusCreated, a)
		}

	}
}

func (s *Server) GetFormsByAuthorUUID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerparts := strings.Split(r.Header.Get("Authorization"), " ")
		name, code, err := Auth_GetAttrs(fmt.Sprintf("http://authenticate:7000/api/v1/check?token=%s", headerparts[1]))
		if err != nil {
			errJSON(w, code, err)
		}

		f, err := s.store.Forms().GetAllFormsByAuthorUUID(name)
		if err != nil {
			errJSON(w, http.StatusUnauthorized, errNoForms)
			return
		}

		toJSON(w, http.StatusOK, f)
	}
}

func (s *Server) DeleteFormByFormUuid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Query().Get("form")
		if err := s.store.Forms().DeleteAllFormByUuid(uuid); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func (s *Server) DeleteFieldById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Query().Get("id")
		if err := s.store.Forms().DeleteOneFieldByFieldId(uuid); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func (s *Server) DeleteAnswerById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		id := params.Get("id")
		if err := s.store.Forms().DeleteOneAnswerById(id); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}
