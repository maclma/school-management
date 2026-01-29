package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/services"
	"github.com/maclma/school-management/pkg/utils"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CREATE USER
func (h *UserHandler) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	if err := h.service.Create(&user); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GET ALL USERS
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET USER BY ID
func (h *UserHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid user id"))
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, utils.NotFound(err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}

// UPDATE USER
func (h *UserHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid user id"))
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	user.ID = uint(id)

	if err := h.service.Update(&user); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}
