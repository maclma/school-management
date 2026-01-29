package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/services"
	"github.com/maclma/school-management/pkg/utils"
)

type TeacherHandler struct {
	service *services.TeacherService
}

func NewTeacherHandler(service *services.TeacherService) *TeacherHandler {
	return &TeacherHandler{service: service}
}

func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	if err := h.service.CreateTeacher(&teacher); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, teacher)
}

func (h *TeacherHandler) GetTeachers(c *gin.Context) {
	teachers, err := h.service.GetTeachers()
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	teachers, total, err := h.service.GetPaginatedTeachers(page, limit)
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	totalPages := (total + int64(limit) - 1) / int64(limit)

	c.Header("X-Total-Count", strconv.FormatInt(total, 10))
	c.Header("X-Total-Pages", strconv.FormatInt(totalPages, 10))

	c.JSON(http.StatusOK, gin.H{
		"data":        teachers,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	})
}

func (h *TeacherHandler) GetTeacherByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid teacher ID"))
		return
	}

	teacher, err := h.service.GetTeacherByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, utils.NotFound("teacher not found"))
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid teacher ID"))
		return
	}

	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	teacher.ID = uint(id)
	if err := h.service.UpdateTeacher(&teacher); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid teacher ID"))
		return
	}

	if err := h.service.DeleteTeacher(uint(id)); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "teacher deleted"})
}
