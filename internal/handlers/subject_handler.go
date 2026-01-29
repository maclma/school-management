package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/services"
	"github.com/maclma/school-management/pkg/utils"
)

type SubjectHandler struct {
	service *services.SubjectService
}

func NewSubjectHandler(service *services.SubjectService) *SubjectHandler {
	return &SubjectHandler{service: service}
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	if err := h.service.CreateSubject(&subject); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	subjects, total, err := h.service.GetPaginatedSubjects(page, limit)
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	totalPages := (total + int64(limit) - 1) / int64(limit)
	c.Header("X-Total-Count", strconv.FormatInt(total, 10))
	c.Header("X-Total-Pages", strconv.FormatInt(totalPages, 10))

	c.JSON(http.StatusCreated, gin.H{
		"data":        subjects,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	})
}

func (h *SubjectHandler) GetSubjects(c *gin.Context) {
	subjects, err := h.service.GetSubjects()
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, subjects)
}

func (h *SubjectHandler) GetSubjectByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid subject ID"))
		return
	}

	subject, err := h.service.GetSubjectByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, utils.NotFound("subject not found"))
		return
	}

	c.JSON(http.StatusOK, subject)
}

func (h *SubjectHandler) GetSubjectsByClass(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("classId"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid class ID"))
		return
	}

	subjects, err := h.service.GetSubjectsByClass(uint(classID))
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, subjects)
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid subject ID"))
		return
	}

	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	subject.ID = uint(id)

	if err := h.service.UpdateSubject(&subject); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, subject)
}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid subject ID"))
		return
	}

	if err := h.service.DeleteSubject(uint(id)); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subject deleted"})
}
func (h *SubjectHandler) GetSubjectsByClassID(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("class_id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid class ID"))
		return
	}
	subjects, err := h.service.GetSubjectsByClass(uint(classID))
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	c.JSON(http.StatusOK, subjects)
}
