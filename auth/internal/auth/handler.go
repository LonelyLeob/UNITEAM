package auth

import (
	"authenticate/internal/auth/api"
	"authenticate/internal/auth/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) RegistrationUser_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.RegisterDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, err, http.StatusBadRequest)
			return
		}

		user := models.SetUser(req.Name, req.Password, req.Email)

		if err := user.InputValidation(); err != nil {
			errJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := s.pgstore.User().CreateUser(user); err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		meta := &models.UserMeta{
			Id: user.Id,
			Lv: time.Now().Unix(),
		}
		user.Meta = append(user.Meta, meta)

		if err := meta.ParseUserAgent(r.UserAgent()); err != nil {
			errJSON(w, api.ErrNoInsertMeta, http.StatusUnprocessableEntity)
			return
		}

		mp, err := s.tg.CreatePairToken(user)
		if err != nil {
			errJSON(w, err, http.StatusBadRequest)
			return
		}
		user.Meta[0].Id = user.Id
		user.Meta[0].Refresh = mp.Refresh

		if err := s.pgstore.UserMeta().SetMetadata(meta); err != nil {
			errJSON(w, api.ErrNoInsertMeta, http.StatusUnprocessableEntity)
			return
		}

		s.rstore.SaveToken(mp.Access, user.Id.String(), s.tg.AccessDuration)

		if mp == nil {
			errJSON(w, api.ErrTokenIsOut, http.StatusBadRequest)
			return
		}

		BindJSON(w, mp, http.StatusOK)
	}
}

func (s *Server) AuthenticateUser_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.AuthenticateDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, err, http.StatusBadRequest)
			return
		}

		user, err := s.pgstore.User().GetAndVerificateUser(req.Name, req.Password)
		if err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		meta := &models.UserMeta{
			Id: user.Id,
			Lv: time.Now().Unix(),
		}
		user.Meta = append(user.Meta, meta)

		if err := meta.ParseUserAgent(r.UserAgent()); err != nil {
			errJSON(w, api.ErrNoInsertMeta, http.StatusUnprocessableEntity)
			return
		}

		mp, err := s.tg.CreatePairToken(user)
		if err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		user.Meta[0].Id = user.Id
		user.Meta[0].Refresh = mp.Refresh
		if err := s.pgstore.UserMeta().CheckForEqualEP(meta); err != nil {
			if err := s.pgstore.UserMeta().SetMetadata(meta); err != nil {
				errJSON(w, api.ErrNoInsertMeta, http.StatusUnprocessableEntity)
				return
			}
		}

		s.rstore.SaveToken(mp.Access, user.Id.String(), s.tg.AccessDuration)

		if mp == nil {
			errJSON(w, api.ErrTokenIsOut, http.StatusBadRequest)
			return
		}

		BindJSON(w, mp, http.StatusOK)
	}
}

func (s *Server) UpdateTokenUser_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh := r.Header.Get("X-Refresh")
		uid, err := s.pgstore.UserMeta().GetUUIDByRefresh(refresh)
		if err != nil {
			errJSON(w, err, http.StatusUnauthorized)
			return
		}

		user, err := s.pgstore.User().GetUserByUid(uid)
		if err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		meta := &models.UserMeta{
			Id: user.Id,
			Lv: time.Now().Unix(),
		}
		user.Meta = append(user.Meta, meta)

		if err := meta.ParseUserAgent(r.UserAgent()); err != nil {
			errJSON(w, api.ErrNoInsertMeta, http.StatusUnprocessableEntity)
			return
		}

		mp, err := s.tg.CreatePairToken(user)
		if err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		user.Meta[0].Id = user.Id
		user.Meta[0].Refresh = mp.Refresh

		if err := s.pgstore.UserMeta().CheckForEqualEP(meta); err != nil {
			if err := s.pgstore.UserMeta().SetMetadata(meta); err != nil {
				errJSON(w, api.ErrNoInsertMeta, http.StatusUnprocessableEntity)
				return
			}
		} else {
			if err := s.pgstore.UserMeta().ResetEP(meta); err != nil {
				errJSON(w, api.ErrNoResetMeta, http.StatusUnprocessableEntity)
				return
			}
		}

		s.rstore.SaveToken(mp.Access, user.Id.String(), s.tg.AccessDuration)

		if mp == nil {
			errJSON(w, api.ErrTokenIsOut, http.StatusBadRequest)
			return
		}

		BindJSON(w, mp, http.StatusOK)
	}
}

func (s *Server) ForgetPassword_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.ForgetPasswordDTO{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, err, http.StatusBadRequest)
			return
		}

		email, err := s.pgstore.User().GetEmailByName(req.Name)
		if err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		if err := s.mailClient.SendMail(
			"Change your password\n",
			fmt.Sprintf(
				"This message was sended by UNITEAM for update your password\nPlease, verify this action on: http://uni-team-inc.online:4000/api/v1/%s?n=%s&p=%s",
				UpdatePwd_Route,
				req.Name,
				req.New,
			),
			email); err != nil {
			errJSON(w, err, http.StatusBadRequest)
			return
		}

		BindJSON(w, "ok", http.StatusOK)
	}
}

func (s *Server) UpdatePassword_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("n")
		pwd := r.URL.Query().Get("p")

		user := &models.User{
			Name: name,
			Pwd:  pwd,
		}

		if err := s.pgstore.User().UpdatePasswordByName(user); err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		BindJSON(w, "ok", http.StatusOK)
	}
}

func (s *Server) DeleteUser_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("n")
		pwd := r.URL.Query().Get("p")
		_, err := s.pgstore.User().GetAndVerificateUser(name, pwd)
		if err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		if err := s.pgstore.User().DeleteUserByName(name); err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		BindJSON(w, "ok", http.StatusOK)
	}
}

func (s *Server) GetUser_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		name, err := s.tg.ParseAccess(bearer)
		if err != nil {
			errJSON(w, err, http.StatusUnauthorized)
			return
		}

		user, err := s.pgstore.User().GetUserByName(name)
		if err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}

		meta, err := s.pgstore.UserMeta().GetMetadataById(user.Id)
		if err != nil {
			errJSON(w, err, http.StatusBadRequest)
			return
		}

		user.Meta = meta

		BindJSON(w, user, http.StatusOK)
	}
}

func (s *Server) LogoutUser_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := s.tg.ParseAccess(r.Header.Get("Authorization"))
		if err == nil {
			ta, err := s.tg.ParseHeader(r.Header.Get("Authorization"))
			if err != nil {
				errJSON(w, err, http.StatusBadRequest)
				return
			}
			if err := s.rstore.DeleteToken(ta); err != nil {
				errJSON(w, err, http.StatusBadRequest)
				return
			}
		}

		tr, err := s.tg.ParseHeader(r.Header.Get("X-Refresh"))
		if err != nil {
			errJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := s.pgstore.UserMeta().DeleteEP(tr); err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}
	}
}

func (s *Server) CloseSession_Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh := r.Header.Get("X-Refresh")
		if err := s.pgstore.UserMeta().DeleteEP(refresh); err != nil {
			errJSON(w, err, http.StatusUnprocessableEntity)
			return
		}
		BindJSON(w, "ok", http.StatusOK)
	}
}
