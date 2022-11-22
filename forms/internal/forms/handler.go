package forms

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) FormsCreatingHttp() http.HandlerFunc {
	type Request struct {
		Name   string `json:"name"`
		Desc   string `json:"desc"`
		Anonym bool   `json:"anon"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		token, err := checkAuth(r.Header.Get("Authorization"))
		if err != nil {
			errJSON(w, http.StatusUnauthorized, err)
			return
		}

		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		author, err := ParseToken(token, s.signingKey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, errBadToParseToken)
			return
		}

		fuid := uuid.New()
		f := &Form{
			Uuid:        fuid,
			Name:        req.Name,
			Description: req.Desc,
			IsAnonym:    req.Anonym,
			AuthorName:  author,
		}

		if err := s.store.Forms().CreateForm(f); err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		} else {
			toJSON(w, http.StatusCreated, f)
		}
	}
}

func (s *Server) FieldCreatingForm() http.HandlerFunc {
	type request struct {
		QuizName string `json:"quiz"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		token, err := checkAuth(r.Header.Get("Authorization"))
		if err != nil {
			errJSON(w, http.StatusUnauthorized, err)
			return
		}
		_, err = ParseToken(token, s.signingKey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, errBadToParseToken)
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		params := r.URL.Query()
		form := params.Get("form")

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
		token, err := checkAuth(r.Header.Get("Authorization"))
		if err != nil {
			errJSON(w, http.StatusUnauthorized, err)
			return
		}
		_, err = ParseToken(token, s.signingKey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, errBadToParseToken)
			return
		}

		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		params := r.URL.Query()
		field := params.Get("field")

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
		token, err := checkAuth(r.Header.Get("Authorization"))
		if err != nil {
			errJSON(w, http.StatusUnauthorized, err)
			return
		}

		author, err := ParseToken(token, s.signingKey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, errBadToParseToken)
			return
		}

		forms, err := s.store.Forms().GetAllFormsByAuthorUUID(author)
		if err != nil {
			errJSON(w, http.StatusUnauthorized, errNoForms)
			return
		}

		toJSON(w, http.StatusOK, forms)
	}
}

func (s *Server) DeleteFormByFormUuid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Allow-Access-Control-Origin", "*")
			return
		}

		params := r.URL.Query()
		uuid := params.Get("form")
		if err := s.store.Forms().DeleteAllFormByUuid(uuid); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func (s *Server) DeleteFieldById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Allow-Access-Control-Origin", "*")
			return
		}

		params := r.URL.Query()
		uuid := params.Get("id")
		if err := s.store.Forms().DeleteOneFieldByFieldId(uuid); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func (s *Server) DeleteAnswerById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Allow-Access-Control-Origin", "*")
			return
		}

		params := r.URL.Query()
		id := params.Get("id")
		if err := s.store.Forms().DeleteOneAnswerById(id); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}
