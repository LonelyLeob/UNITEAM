package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (s *server) SendMessageHandler() http.HandlerFunc {
	type payload struct {
		Text string `json:"text"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		pl := &payload{}
		if err := json.NewDecoder(r.Body).Decode(pl); err != nil {
			return
		}

		msg := &Message{
			CreatedAt: time.Now().Unix(),
			Text:      pl.Text,
		}

		token, err := GetAuthString(r.Header.Get("Authorization"))
		switch err {
		case nil:
			msg.From = token
		default:
			msg.From = "Anonymous"
		}

		if err := s.mongoose.CreateMessage(msg); err != nil {
			return
		}
	}
}

func (s *server) GetBatchMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lim := r.URL.Query().Get("limit")
		limit, err := strconv.ParseInt(lim, 10, 64)
		if err != nil {
			return
		}
		msgs, err := s.mongoose.GetMessagesWithOffset(limit)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(msgs)
	}
}
