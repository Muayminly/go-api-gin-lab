// 6609650491
// Piyatida Reakdee
package repositories

import (
	"database/sql"

	"example.com/student-api/models"
)

type StudentRepository struct {
	DB *sql.DB
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.DB.Query("SELECT id, name, major, gpa FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		if err := rows.Scan(&s.Id, &s.Name, &s.Major, &s.GPA); err != nil {
			return nil, err
		}

		students = append(students, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepository) GetByID(id string) (*models.Student, error) {
	row := r.DB.QueryRow(
		"SELECT id, name, major, gpa FROM students WHERE id = ?",
		id,
	)

	var s models.Student
	if err := row.Scan(&s.Id, &s.Name, &s.Major, &s.GPA); err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepository) Create(s models.Student) error {
	_, err := r.DB.Exec(
		"INSERT INTO students (id, name, major, gpa) VALUES (?, ?, ?, ?)",
		s.Id, s.Name, s.Major, s.GPA,
	)
	return err
}

func (r *StudentRepository) UpdateByID(id string, name string, major string, gpa float64) (bool, error) {
	res, err := r.DB.Exec(
		"UPDATE students SET name = ?, major = ?, gpa = ? WHERE id = ?",
		name, major, gpa, id,
	)
	if err != nil {
		return false, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func (r *StudentRepository) DeleteByID(id string) (bool, error) {
	res, err := r.DB.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return false, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}
