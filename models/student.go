// 6609650491
// Piyatida Reakdee
package models

import "strings"

type Student struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Major string  `json:"major"`
	GPA   float64 `json:"gpa"`
}

func (s Student) ValidateCreate() (bool, string) {
	if strings.TrimSpace(s.Id) == "" {
		return false, "ID must not be empty"
	}
	if strings.TrimSpace(s.Name) == "" {
		return false, "Name must not be empty"
	}
	if s.GPA < 0.0 || s.GPA > 4.0 {
		return false, "GPA must be between 0.00 and 4.00"
	}
	return true, ""
}

func (s Student) ValidateUpdate() (bool, string) {
	if strings.TrimSpace(s.Name) == "" {
		return false, "Name must not be empty"
	}
	if s.GPA < 0.0 || s.GPA > 4.0 {
		return false, "GPA must be between 0.00 and 4.00"
	}
	return true, ""
}
