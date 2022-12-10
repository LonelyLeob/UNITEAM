package store

import (
	"authenticate/internal/auth/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	// debug = os.Getenv("DEBUG")
	debug = "True"
)

const (
	UserTable = "users"

	createUserQuery     = "INSERT INTO %s (id, name, email, password) VALUES ($1, $2, $3, $4)"
	getEmailByNameQuery = "SELECT email FROM %s WHERE name = $1"
	getUserByNameQuery  = "SELECT id, email, password FROM %s WHERE name = $1"
	updatePasswordQuery = "UPDATE %s SET password = $1 WHERE name = $2"
	deleteUserQuery     = "DELETE FROM %s WHERE name = $1"
)

type UserRepository struct {
	ps *PostgresStore
}

func (u *UserRepository) CreateUser(user *models.User) error {
	var err error
	if user == nil {
		return errNilPtr
	}

	user.Id = uuid.New()
	user.HashPassword(user.Pwd)

	if _, err = u.ps.db.Exec(
		fmt.Sprintf(createUserQuery, UserTable),
		user.Id,
		user.Name,
		user.Email,
		user.Pwd,
	); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errUnreachableAction
	}

	return nil
}

func (u *UserRepository) GetEmailByName(name string) (string, error) {
	var email string
	if err := u.ps.db.QueryRow(
		fmt.Sprintf(getEmailByNameQuery, UserTable),
		name,
	).Scan(
		&email,
	); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return "", errUnreachableAction
	}

	return email, nil
}

func (u *UserRepository) GetUserByName(name string) (*models.User, error) {
	user := &models.User{}
	if err := u.ps.db.QueryRow(
		fmt.Sprintf(getUserByNameQuery, UserTable),
		name,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Pwd,
	); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return nil, errUnreachableAction
	}

	user.Name = name
	return user, nil
}

func (u *UserRepository) GetAndVerificateUser(name, pwd string) (*models.User, error) {
	user := &models.User{}
	if err := u.ps.db.QueryRow(
		fmt.Sprintf(getUserByNameQuery, UserTable),
		name,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Pwd,
	); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return nil, errUnreachableAction
	}

	if err := user.ComparePassword(pwd); err != nil {
		return nil, err
	}

	user.Meta = []*models.UserMeta{}
	user.Name = name
	return user, nil
}

func (u *UserRepository) DeleteUserById(name string) error {
	if _, err := u.ps.db.Exec(
		fmt.Sprintf(deleteUserQuery, UserTable),
		name,
	); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errUnreachableAction
	}

	return nil
}

func (u *UserRepository) UpdatePasswordByName(user *models.User) error {
	user.HashPassword(user.Pwd)

	if _, err := u.ps.db.Exec(
		fmt.Sprintf(updatePasswordQuery, UserTable),
		user.Pwd,
		user.Name,
	); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errUnreachableAction
	}

	return nil
}
