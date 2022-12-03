package forms

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) FormsCreatingHttp() http.HandlerFunc {
	type ReqFormsDTO struct {
		Name   string `json:"name"`
		Desc   string `json:"desc"`
		Anonym bool   `json:"anon"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		freq := &ReqFormsDTO{}
		if err := json.NewDecoder(r.Body).Decode(freq); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		name, err := ParseTokenFromHeader(r.Header.Get("Authorization"), s.signingKey)
		if err != nil {
			errJSON(w, http.StatusUnauthorized, err)
			return
		}

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
	}

}

func (s *Server) FieldCreatingForm() http.HandlerFunc {
	type request struct {
		QuizName string `json:"quiz"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
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
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		a := &VariableAnswers{
			Answer: req.Answer,
		}

		if err := a.ParseFieldId(r.URL.Query().Get("field")); err != nil {
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
		name, err := ParseTokenFromHeader(r.Header.Get("Authorization"), s.signingKey)
		if err != nil {
			return
		}

		fs, err := s.store.Forms().GetAllFormsByAuthorUUID(name)
		if err != nil {
			errJSON(w, http.StatusUnprocessableEntity, errNoForms)
			return
		}

		toJSON(w, http.StatusOK, fs)
	}
}

func (s *Server) DeleteFormByFormUuid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Query().Get("form")
		if err := s.store.Forms().DeleteAllFormByUuid(uuid); err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func (s *Server) DeleteFieldById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Query().Get("id")
		if err := s.store.Forms().DeleteOneFieldByFieldId(uuid); err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func (s *Server) DeleteAnswerById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if err := s.store.Forms().DeleteOneAnswerById(id); err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}
