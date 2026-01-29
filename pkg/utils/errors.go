package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppError represents a structured error
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// RespondWithError sends a standard JSON error response
func RespondWithError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// Predefined error helpers
func BadRequest(msg string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func Unauthorized(msg string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func Forbidden(msg string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}

func NotFound(msg string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func Internal(msg string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
