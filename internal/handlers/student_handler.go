package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/services"
	"github.com/maclma/school-management/pkg/utils"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	if err := h.service.CreateStudent(&student); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, student)
}
func (h *StudentHandler) GetAll(c *gin.Context) {
	students, err := h.service.GetAll()
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	c.JSON(http.StatusOK, students)
}
func (h *StudentHandler) GetStudents(c *gin.Context) {
	// Parse query params
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		utils.RespondWithError(c, utils.BadRequest("invalid page number"))
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		utils.RespondWithError(c, utils.BadRequest("invalid limit number"))
		return
	}

	students, total, err := h.service.GetPaginated(page, limit)
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  students,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid student ID"))
		return
	}

	student, err := h.service.GetStudentByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, utils.NotFound("student not found"))
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid student ID"))
		return
	}

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	student.ID = uint(id)

	if err := h.service.UpdateStudent(&student); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid student ID"))
		return
	}

	if err := h.service.DeleteStudent(uint(id)); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student deleted"})
}
