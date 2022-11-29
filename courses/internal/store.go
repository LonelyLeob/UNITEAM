package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	config *ConfigDatabase
	db     *sql.DB
}

func NewStorage(config *ConfigDatabase) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Connect() error {
	fmt.Println(s.config.CreateFullDBAddr())
	db, err := sql.Open("postgres", s.config.CreateFullDBAddr())
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

// add new course
func (s *Storage) CreateCourse(course *Course) error {
	if err := s.db.QueryRow(
		"INSERT INTO courses (title, cdesc, author) VALUES ($1, $2, $3) RETURNING id",
		course.Title,
		course.Desc,
		course.Author,
	).Scan(&course.CourseId); err != nil {
		return err
	}

	return nil
}

// add new section of concrete course
func (s *Storage) addSectionCourse(section *CourseSection) error {
	if err := s.db.QueryRow(
		"INSERT INTO sections (course_id, content) VALUES ($1, $2) RETURNING id",
		section.Course,
		section.Content,
	).Scan(&section.SectionId); err != nil {
		return err
	}

	return nil
}

// get shorthand data for mainpage
func (s *Storage) SelectShortDataCourses() ([]*Course, error) {
	var crs []*Course
	rows, err := s.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := Course{}
		if err := rows.Scan(&c.CourseId, &c.Title, &c.Desc, &c.Author); err != nil {
			fmt.Println(err)
			continue
		}

		crs = append(crs, &c)
	}

	return crs, nil
}

func (s *Storage) SelectCourseDataById(id int) (*Course, error) {
	course := &Course{}
	if err := s.db.QueryRow(
		"SELECT id, title, cdesc, author FROM courses WHERE id = $1",
		id,
	).Scan(
		&course.CourseId,
		&course.Title,
		&course.Desc,
		&course.Author,
	); err != nil {
		return nil, err
	}
	var arrS []*CourseSection

	rows, err := s.db.Query("SELECT id, course_id, content FROM sections WHERE course_id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		sect := &CourseSection{}
		if err := rows.Scan(&sect.SectionId, &sect.Course, &sect.Content); err != nil {
			fmt.Println(err)
			continue
		}

		arrS = append(arrS, sect)
	}

	course.Sections = arrS
	return course, nil
}

// delete full course if u need
func (s *Storage) DeleteCourseById(id int) error {
	if _, err := s.db.Exec("DELETE FROM sections WHERE course_id = $1", id); err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM courses WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}

// delete full section if u need
func (s *Storage) DeleteSectionById(id int) error {
	if _, err := s.db.Exec("DELETE FROM sections WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateSectionContent(id int, content string) error {
	if _, err := s.db.Exec("UPDATE sections SET content = $1 WHERE id = $2", content, id); err != nil {
		return err
	}

	return nil
}
