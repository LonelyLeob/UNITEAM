package internal

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Repos interface {
	GetUser(*User, string) (*User, error)
	CreateUser(*User) error
}

type repos struct {
	store *store
}

func (r *repos) GetUser(user *User, pwd string) (*User, error) {
	if err := r.store.conn.QueryRow(
		"SELECT name::TEXT, email::TEXT, role::TEXT, password::TEXT FROM users WHERE name = $1;",
		user.name,
	).Scan(
		&user.name,
		&user.email,
		&user.role,
		&user.password,
	); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.password),
		[]byte(pwd),
	); err != nil {
		return nil, err
	}

	r.updateVisit(user)

	return user, nil
}

func (r *repos) CreateUser(user *User) error {
	user.HashPassword(user.password)

	if err := r.store.conn.QueryRow(
		"INSERT INTO users (name, password, role, email, lastv) VALUES ($1, $2::VARCHAR, $3, $4, $5::BIGINT) RETURNING name;",
		user.name,
		user.password,
		user.role,
		user.email,
		time.Now().Unix(),
	).Scan(&user.name); err != nil {
		return err
	}

	return nil
}

func (r *repos) updateVisit(user *User) error {
	if err := r.store.conn.QueryRow(
		"UPDATE users SET lastv = $1 WHERE name = $2",
		time.Now().Unix(),
		user.name,
	).Scan(); err != nil {
		return err
	}

	return nil
}
