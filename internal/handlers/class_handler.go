package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/services"
	"github.com/maclma/school-management/pkg/utils"
)

type ClassHandler struct {
	service *services.ClassService
}

func NewClassHandler(service *services.ClassService) *ClassHandler {
	return &ClassHandler{service: service}
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	if err := h.service.CreateClass(&class); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, class)
}

func (h *ClassHandler) GetClasses(c *gin.Context) {
	classes, err := h.service.GetClasses()
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, classes)
}

func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid class ID"))
		return
	}

	class, err := h.service.GetClassByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, utils.NotFound("class not found"))
		return
	}

	c.JSON(http.StatusOK, class)
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid class ID"))
		return
	}

	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	class.ID = uint(id)
	if err := h.service.UpdateClass(&class); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, class)
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid class ID"))
		return
	}

	if err := h.service.DeleteClass(uint(id)); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "class deleted"})
}
