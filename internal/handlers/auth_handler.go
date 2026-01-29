package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/services"
	"github.com/maclma/school-management/pkg/utils"
)

// LoginRequest represents login payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthHandler handles authentication endpoints
type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login endpoint
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	user, token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		utils.RespondWithError(c, utils.Unauthorized(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

// Me endpoint (returns current user)
func (h *AuthHandler) Me(c *gin.Context) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		utils.RespondWithError(c, utils.Unauthorized("user not found in context"))
		return
	}

	userID, _ := strconv.Atoi(userIDStr.(string))
	user, err := h.service.GetUserByID(uint(userID))
	if err != nil {
		utils.RespondWithError(c, utils.NotFound("user not found"))
		return
	}

	c.JSON(http.StatusOK, user)
}
