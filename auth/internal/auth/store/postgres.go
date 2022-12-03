package store

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	addr   string
	db     *sql.DB
	urepo  *UserRepository
	umrepo *UserMetaRepository
}

func NewStore(addr string) *PostgresStore {
	return &PostgresStore{
		addr: addr,
	}
}

func (p *PostgresStore) InitConnect(ctx context.Context) error {
	db, err := sql.Open("postgres", p.addr)
	if err != nil {
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		return err
	}
	p.db = db

	return nil
}

func (p *PostgresStore) Disconnect() error {
	if err := p.db.Close(); err != nil {
		return err
	}

	return nil
}

func (p *PostgresStore) User() *UserRepository {
	if p.urepo != nil {
		return p.urepo
	}

	p.urepo = &UserRepository{
		ps: p,
	}

	return p.urepo
}

func (p *PostgresStore) UserMeta() *UserMetaRepository {
	if p.umrepo != nil {
		return p.umrepo
	}

	p.umrepo = &UserMetaRepository{
		ps: p,
	}

	return p.umrepo
}
