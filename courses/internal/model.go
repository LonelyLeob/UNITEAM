package internal

import "github.com/golang-jwt/jwt/v4"

type Course struct {
	CourseId int              `json:"id"`
	Title    string           `json:"title"`
	Desc     string           `json:"desc"`
	Author   string           `json:"author"`
	Sections []*CourseSection `json:"sections,omitempty"`
}

type CourseSection struct {
	SectionId int    `json:"id"`
	Course    int    `json:"course_id"`
	Content   string `json:"content"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	Name string `json:"name"`
}
