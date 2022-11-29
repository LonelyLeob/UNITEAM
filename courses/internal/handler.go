package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// get courses shortcut data (id, title, short desc, price)
func (s *server) getShorthandCourses_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses, err := s.store.SelectShortDataCourses()
		if err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, courses)
	}
}

// get full course by id param
func (s *server) getCourseData_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cid, err := strconv.Atoi(r.URL.Query().Get("course"))
		if err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		course, err := s.store.SelectCourseDataById(cid)
		if err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, course)
	}
}

// add course with struct params
func (s *server) addCourse_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author, err := ParseTokenFromHeader(r.Header.Get("Authorization"), []byte(s.config.SigningKey))
		if err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		c := &Course{
			Author: author,
		}

		if err := json.NewDecoder(r.Body).Decode(c); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		if err := s.store.CreateCourse(c); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusCreated, c)
	}
}

// update course data by id param
func (s *server) updateSectionContent_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sect := &CourseSection{}
		if err := json.NewDecoder(r.Body).Decode(sect); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}
		cid, err := strconv.Atoi(r.URL.Query().Get("section"))
		if err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}
		if err := s.store.UpdateSectionContent(cid, sect.Content); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"upd_content": sect.Content})
	}
}

// delete course by id param
func (s *server) deleteCourse_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cid, err := strconv.Atoi(r.URL.Query().Get("course"))
		if err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		if err := s.store.DeleteCourseById(cid); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func (s *server) addSectionCourse_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sect := &CourseSection{}
		if err := json.NewDecoder(r.Body).Decode(sect); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		if err := s.store.addSectionCourse(sect); err != nil {
			errJSON(w, http.StatusInternalServerError, err)
			return
		}

		toJSON(w, http.StatusCreated, sect)
	}
}
