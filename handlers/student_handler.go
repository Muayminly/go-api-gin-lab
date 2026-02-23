// 6609650491
// Piyatida Reakdee
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func jsonError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"error": msg})
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		// service จะส่ง internal มาเป็น ServiceError อยู่แล้ว แต่กันพลาดไว้
		jsonError(c, http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")

	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		if se, ok := err.(*services.ServiceError); ok {
			switch se.Kind {
			case services.ErrNotFound:
				jsonError(c, http.StatusNotFound, se.Message)
				return
			case services.ErrValidation:
				jsonError(c, http.StatusBadRequest, se.Message)
				return
			case services.ErrConflict:
				jsonError(c, http.StatusBadRequest, se.Message)
				return
			default:
				jsonError(c, http.StatusInternalServerError, "Internal server error")
				return
			}
		}
		jsonError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		jsonError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	err := h.Service.CreateStudent(student)
	if err != nil {
		if se, ok := err.(*services.ServiceError); ok {
			switch se.Kind {
			case services.ErrValidation:
				jsonError(c, http.StatusBadRequest, se.Message)
				return
			case services.ErrConflict:
				jsonError(c, http.StatusBadRequest, se.Message)
				return
			default:
				jsonError(c, http.StatusInternalServerError, "Internal server error")
				return
			}
		}
		jsonError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var payload models.Student
	if err := c.ShouldBindJSON(&payload); err != nil {
		jsonError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	updated, err := h.Service.UpdateStudent(id, payload)
	if err != nil {
		if se, ok := err.(*services.ServiceError); ok {
			switch se.Kind {
			case services.ErrValidation:
				jsonError(c, http.StatusBadRequest, se.Message)
				return
			case services.ErrNotFound:
				jsonError(c, http.StatusNotFound, se.Message)
				return
			default:
				jsonError(c, http.StatusInternalServerError, "Internal server error")
				return
			}
		}
		jsonError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteStudent(id)
	if err != nil {
		if se, ok := err.(*services.ServiceError); ok {
			switch se.Kind {
			case services.ErrNotFound:
				jsonError(c, http.StatusNotFound, se.Message)
				return
			default:
				jsonError(c, http.StatusInternalServerError, "Internal server error")
				return
			}
		}
		jsonError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.Status(http.StatusNoContent)
}
