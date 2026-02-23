// 6609650491
// Piyatida Reakdee
package services

import (
	"database/sql"
	"errors"
	"strings"

	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

var (
	ErrNotFound   = errors.New("Student not found")
	ErrValidation = errors.New("Validation error")
	ErrInternal   = errors.New("Internal server error")
	ErrConflict   = errors.New("Conflict error")
)

type ServiceError struct {
	Kind    error
	Message string
}

func (e *ServiceError) Error() string { return e.Message }

func notFound() error             { return &ServiceError{Kind: ErrNotFound, Message: "Student not found"} }
func internal() error             { return &ServiceError{Kind: ErrInternal, Message: "Internal server error"} }
func validation(msg string) error { return &ServiceError{Kind: ErrValidation, Message: msg} }
func conflict(msg string) error   { return &ServiceError{Kind: ErrConflict, Message: msg} }

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	students, err := s.Repo.GetAll()
	if err != nil {
		return nil, internal()
	}
	return students, nil
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	stu, err := s.Repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, notFound()
		}
		return nil, internal()
	}
	return stu, nil
}

func (s *StudentService) CreateStudent(student models.Student) error {
	if ok, msg := student.ValidateCreate(); !ok {
		return validation(msg)
	}

	err := s.Repo.Create(student)
	if err != nil {
		// ไม่ expose raw DB error
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return conflict("id already exists")
		}
		return internal()
	}
	return nil
}

func (s *StudentService) UpdateStudent(id string, payload models.Student) (*models.Student, error) {
	if ok, msg := payload.ValidateUpdate(); !ok {
		return nil, validation(msg)
	}

	found, err := s.Repo.UpdateByID(id, payload.Name, payload.Major, payload.GPA)
	if err != nil {
		return nil, internal()
	}
	if !found {
		return nil, notFound()
	}

	updated := &models.Student{
		Id:    id,
		Name:  payload.Name,
		Major: payload.Major,
		GPA:   payload.GPA,
	}
	return updated, nil
}

func (s *StudentService) DeleteStudent(id string) error {
	found, err := s.Repo.DeleteByID(id)
	if err != nil {
		return internal()
	}
	if !found {
		return notFound()
	}
	return nil
}
