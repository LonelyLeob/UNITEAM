package internal

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Storage interface {
	Connect(string)
	Web() Repos
}

func NewStore() Storage {
	return &store{
		maxAttempts: 5,
	}
}

type store struct {
	conn        *sql.DB
	repos       Repos
	maxAttempts int
}

func (s *store) Connect(url string) {
	conn, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < s.maxAttempts; i++ {
		err := conn.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Second * 3)
	}

	s.conn = conn
}

func (s *store) Web() Repos {
	if s.repos != nil {
		return s.repos
	}

	s.repos = &repos{
		store: s,
	}

	return s.repos
}
