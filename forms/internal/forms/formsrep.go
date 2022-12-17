package forms

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type FormsRepository struct {
	store *Store
}

func (r *FormsRepository) CreateForm(data *Form) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO form (uuid, quizName, quizDesc, anonym, authorName) VALUES ($1, $2, $3, $4, $5) RETURNING uuid",
		data.Uuid,
		data.Name,
		data.Description,
		data.IsAnonym,
		data.AuthorName,
	).Scan(
		&data.Uuid,
	); err != nil {
		return err
	}
	return nil
}

func (r *FormsRepository) AddField(data *Field) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO field (formUuid, fieldName) VALUES ($1, $2) RETURNING id",
		&data.FormUuid,
		&data.Quiz,
	).Scan(
		&data.Id,
	); err != nil {
		return err
	}

	return nil
}

func (r *FormsRepository) AddAnswer(data *VariableAnswers) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO answer (answer, fieldId) VALUES ($1, $2) RETURNING id",
		&data.Answer,
		&data.FieldId,
	).Scan(
		&data.Id,
	); err != nil {
		return err
	}

	return nil
}

func (r *FormsRepository) GetFormByUUIDParams(form *Form, stringUuid string) error {
	normalUuid, err := uuid.Parse(stringUuid)
	if err != nil {
		return errNotCorrectUUID
	}
	if err := r.store.db.QueryRow(
		"SELECT * FROM form WHERE uuid = $1",
		normalUuid,
	).Scan(
		&form.Uuid,
		&form.Name,
		&form.Description,
		&form.IsAnonym,
		&form.AuthorName,
	); err != nil {
		return errCantShowForm
	}

	form.Fields, _ = r.GetFieldsWithAnswersByUUID(&form.Uuid)

	return nil
}

func (r *FormsRepository) GetFieldsWithAnswersByUUID(formUuid *uuid.UUID) ([]*Field, error) {
	rows, err := r.store.db.Query("SELECT * FROM field WHERE formUuid = $1", formUuid)
	if err != nil {
		return nil, err
	}

	var arrF []*Field

	for rows.Next() {
		var f Field
		if err := rows.Scan(&f.Id, &f.FormUuid, &f.Quiz); err != nil {
			fmt.Println(err)
			continue
		}
		f.Answers, _ = r.GetAnswersByFieldId(f.Id)
		arrF = append(arrF, &f)
	}
	return arrF, nil
}

func (r *FormsRepository) GetAnswersByFieldId(id uint64) ([]*VariableAnswers, error) {
	rows, err := r.store.db.Query("SELECT * FROM answer WHERE fieldid = $1", id)
	if err != nil {
		return nil, err
	}

	var arrA []*VariableAnswers

	for rows.Next() {
		var a VariableAnswers
		if err := rows.Scan(&a.Id, &a.Answer, &a.FieldId); err != nil {
			fmt.Println(err)
			continue
		}
		arrA = append(arrA, &a)
	}
	return arrA, nil
}

func (r *FormsRepository) GetAllFormsByAuthor(author string) ([]*Form, error) {
	rows, err := r.store.db.Query("SELECT * FROM form WHERE authorname = $1", author)
	if err != nil {
		return nil, err
	}

	var arrF []*Form

	for rows.Next() {
		var f Form
		if err := rows.Scan(&f.Uuid, &f.Name, &f.Description, &f.IsAnonym, &f.AuthorName); err != nil {
			fmt.Println(err)
			continue
		}

		arrF = append(arrF, &f)
	}

	return arrF, nil
}

func (r *FormsRepository) GetFullFormByUUID(uid string) (*Form, error) {
	var f Form
	var err error
	uidp, err := uuid.Parse(uid)
	if err != nil {
		return nil, err
	}

	if err = r.store.db.QueryRow("SELECT * FROM form WHERE uuid = $1", uidp).Scan(
		&f.Uuid, &f.Name, &f.Description, &f.IsAnonym, &f.AuthorName,
	); err != nil {
		return nil, err
	}

	fields, err := r.GetFieldsWithAnswersByUUID(&uidp)
	if err != nil {
		return nil, err
	}

	f.Fields = fields

	return &f, nil
}

func (r *FormsRepository) DeleteAllFormByUuid(struuid string) error {
	normalUuid, err := uuid.Parse(struuid)
	if err != nil {
		return err
	}

	if err := r.store.db.QueryRow("SELECT uuid FROM form WHERE uuid = $1", normalUuid).Scan(&normalUuid); err != nil {
		return err
	}

	if err := r.DeleteAllFieldsByFormUuid(struuid); err != nil {
		return err
	}

	_, err = r.store.db.Exec("DELETE FROM form WHERE uuid=$1", normalUuid)
	if err != nil {
		return err
	}

	return nil
}

func (r *FormsRepository) DeleteAllFieldsByFormUuid(fuid string) error {
	uid, err := uuid.Parse(fuid)
	if err != nil {
		return err
	}

	rows, _ := r.store.db.Query("SELECT id FROM field WHERE formUuid = $1", uid)

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			continue
		}
		strid := strconv.Itoa(id)
		fmt.Println(id)
		r.DeleteAllAnswersByFieldId(strid)
	}

	_, err = r.store.db.Exec("DELETE FROM field WHERE formUuid = $1", uid)
	if err != nil {
		return err
	}

	return nil
}

func (r *FormsRepository) DeleteAllAnswersByFieldId(strid string) error {
	id, err := strconv.Atoi(strid)
	if err != nil {
		return err
	}

	_, err = r.store.db.Exec("DELETE FROM answer WHERE fieldId = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *FormsRepository) DeleteOneFieldByFieldId(strid string) error {
	id, err := strconv.Atoi(strid)
	if err != nil {
		return err
	}

	if err := r.store.db.QueryRow("SELECT id from field WHERE id = $1", id).Scan(&id); err != nil {
		return err
	}

	if err := r.DeleteAllAnswersByFieldId(strid); err != nil {
		return err
	}

	_, err = r.store.db.Exec("DELETE FROM field WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *FormsRepository) DeleteOneAnswerById(strid string) error {
	id, err := strconv.Atoi(strid)
	if err != nil {
		return err
	}

	if err := r.store.db.QueryRow("SELECT id from answer WHERE id = $1", id).Scan(&id); err != nil {
		return err
	}

	_, err = r.store.db.Exec("DELETE FROM answer WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
