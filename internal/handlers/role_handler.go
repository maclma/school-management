package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/services"
	"github.com/maclma/school-management/pkg/utils"
)

type RoleHandler struct {
	service *services.RoleService
}

func NewRoleHandler(service *services.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	if err := h.service.CreateRole(&role); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) GetRoles(c *gin.Context) {
	roles, err := h.service.GetRoles()
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid role ID"))
		return
	}

	role, err := h.service.GetRoleByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, utils.NotFound("role not found"))
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid role ID"))
		return
	}

	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	role.ID = uint(id)
	if err := h.service.UpdateRole(&role); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid role ID"))
		return
	}

	if err := h.service.DeleteRole(uint(id)); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role deleted"})
}
