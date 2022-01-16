package dto

import (
	"lms-api/internal/model"
	res "lms-api/pkg/util/response"
)

// var resultCourse []model.CourseEntityModel

type CourseGetResponse2 struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Kode        string `json:"kode"`
	Certificate bool   `json:"certificate"`
	Thumbnail   string `json:"thumbnail"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Price       int    `json:"price"`
	Level       string `json:"level"`
	Description string `json:"description"`
	MentorID    int    `json:"mentor_id"`
	CategoryID  int    `json:"category_id"`
	Mentor      struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Profile    string `json:"profile"`
		Profession string `json:"profession"`
		Email      string `json:"email"`
	} `json:"mentor"`
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
}

func FormatCourse(course model.CourseEntityModel) CourseGetResponse2 {
	formatter := CourseGetResponse2{
		ID:          course.ID,
		Name:        course.Name,
		Kode:        course.Kode,
		Certificate: course.Certificate,
		Thumbnail:   course.Thumbnail,
		Type:        string(course.Type),
		Status:      string(course.Status),
		Price:       course.Price,
		Level:       string(course.Level),
		Description: course.Description,
		MentorID:    course.MentorID,
		CategoryID:  course.CategoryID,
		Mentor: struct {
			ID         int    "json:\"id\""
			Name       string "json:\"name\""
			Profile    string "json:\"profile\""
			Profession string "json:\"profession\""
			Email      string "json:\"email\""
		}{
			ID:         course.Mentor.ID,
			Name:       course.Mentor.Name,
			Profile:    course.Mentor.Profile,
			Profession: course.Mentor.Profession,
		},
		Category: struct {
			ID   int    "json:\"id\""
			Name string "json:\"name\""
		}{
			ID:   course.Category.ID,
			Name: course.Category.Name,
		},
	}

	return formatter
}

func FormatCourses(courses []model.CourseEntityModel) []CourseGetResponse2 {

	coursesFormatter := []CourseGetResponse2{}

	for _, course := range courses {
		courseFormatter := FormatCourse(course)
		coursesFormatter = append(coursesFormatter, courseFormatter)
	}

	return coursesFormatter
}

// // Get
type CourseGetResponse struct {
	Datas []model.CourseEntityModel
}

type CourseGetResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data CourseGetResponse `json:"data"`
	} `json:"body"`
}

// GetByID
type CourseGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type CourseGetByIDResponse struct {
	model.CourseEntityModel
}

type CourseGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data CourseGetByIDResponse `json:"data"`
	} `json:"body"`
}

// create
type CourseCreateRequest struct {
	model.CourseEntity
}

type CourseCreateResponse struct {
	model.CourseEntityModel
}

type CourseCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data CourseCreateResponse `json:"data"`
	} `json:"body"`
}

type CourseUpdateRequest struct {
	ID int `param:"id"`
	model.CourseEntity
}

type CourseUpdateResponse struct {
	model.CourseEntityModel
}

type CourseUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data CourseUpdateResponse `json:"data"`
	} `json:"body"`
}

// delete
type CourseDeleteRequest struct {
	ID int `param:"id" validateL:"required,numeric"`
}

type CourseDeleteResponse struct {
	model.CourseEntityModel
}

type CourseDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data CourseDeleteResponse `json:"data"`
	} `json:"body"`
}
