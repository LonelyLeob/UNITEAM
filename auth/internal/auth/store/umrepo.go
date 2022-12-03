package store

import (
	"authenticate/internal/auth/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	MetaTable = "meta"

	userMetaAddQuery     = "INSERT INTO %s (uuid, lv, browser, os, refresh) VALUES ($1, $2, $3, $4, $5)"
	userMetaGetQueryById = "SELECT lv, browser, os FROM %s WHERE uuid = $1"
	userMetaCheckSess    = "select os, browser from %s where os = $1 and browser = $2"
	userDeleteSess       = "DELETE FROM %s WHERE refresh = $1"
)

type UserMetaRepository struct {
	ps *PostgresStore
}

func (um *UserMetaRepository) SetMetadata(meta *models.UserMeta) error {
	if meta == nil {
		return errNilPtr
	}

	if _, err := um.ps.db.Exec(
		fmt.Sprintf(userMetaAddQuery, MetaTable),
		meta.Id,
		meta.Lv,
		meta.Browser,
		meta.OS,
		meta.Refresh,
	); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}
		return errUnreachableAction
	}

	return nil
}

func (um *UserMetaRepository) GetMetadataById(id uuid.UUID) ([]*models.UserMeta, error) {
	meta := []*models.UserMeta{}

	rows, err := um.ps.db.Query(
		fmt.Sprintf(userMetaGetQueryById, MetaTable),
		id,
	)
	if err != nil {
		if debug == "True" {
			logrus.Error(err)
		}
		return nil, errUnreachableAction
	}

	for rows.Next() {
		var sess models.UserMeta
		if err := rows.Scan(&sess.Lv, &sess.Browser, &sess.OS); err != nil {
			fmt.Println(err)
			continue
		}
		sess.Id = id
		meta = append(meta, &sess)
	}

	return meta, nil
}

func (um *UserMetaRepository) CheckForEqualEP(meta *models.UserMeta) error {
	var osdb, brwdb string
	err := um.ps.db.QueryRow(
		fmt.Sprintf(userMetaCheckSess, MetaTable),
		meta.OS,
		meta.Browser,
	).Scan(&osdb, &brwdb)
	fmt.Println(meta.OS, meta.Browser)
	if err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errUnreachableAction
	}
	return nil
}

func (um *UserMetaRepository) DeleteEP(refresh string) error {
	if _, err := um.ps.db.Exec(fmt.Sprintf(userDeleteSess, MetaTable), refresh); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		return errUnreachableAction
	}

	return nil
}
