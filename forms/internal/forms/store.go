package forms

import "database/sql"

type Store struct {
	db  *sql.DB
	rep *FormsRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Forms() *FormsRepository {
	if s.rep != nil {
		return s.rep
	}

	s.rep = &FormsRepository{
		store: s,
	}

	return s.rep
}
