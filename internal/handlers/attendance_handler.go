package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/services"
	"github.com/maclma/school-management/pkg/utils"
)

type AttendanceHandler struct {
	service *services.AttendanceService
}

func NewAttendanceHandler(service *services.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{service: service}
}

// ---------------- CREATE ----------------
func (h *AttendanceHandler) CreateAttendance(c *gin.Context) {
	var attendance models.Attendance

	if err := c.ShouldBindJSON(&attendance); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}

	if err := h.service.Create(&attendance); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	attendances, total, err := h.service.GetPaginatedAttendance(page, limit)
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	totalPages := (total + int64(limit) - 1) / int64(limit)
	c.Header("X-Total-Count", strconv.FormatInt(total, 10))
	c.Header("X-Total-Pages", strconv.FormatInt(totalPages, 10))

	c.JSON(http.StatusCreated, gin.H{
		"data":        attendances,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	})
}

// ---------------- GET ALL ----------------
func (h *AttendanceHandler) GetAllAttendance(c *gin.Context) {
	records, err := h.service.GetAllAttendance()
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, records)
}

// ---------------- GET BY STUDENT ----------------
func (h *AttendanceHandler) GetAttendanceByStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid student id"))
		return
	}

	records, err := h.service.GetAttendanceByStudent(uint(id))
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, records)
}

// ---------------- GET BY CLASS & DATE ----------------
func (h *AttendanceHandler) GetAttendanceByClassAndDate(c *gin.Context) {
	classID, err := strconv.Atoi(c.Query("class_id"))
	if err != nil {
		utils.RespondWithError(c, utils.BadRequest("invalid class id"))
		return
	}

	date := c.Query("date")
	if date == "" {
		utils.RespondWithError(c, utils.BadRequest("date is required"))
		return
	}

	records, err := h.service.GetAttendanceByClassAndDate(uint(classID), date)
	if err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}

	c.JSON(http.StatusOK, records)
}
func (h *AttendanceHandler) MarkAttendance(c *gin.Context) {
	var attendance models.Attendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		utils.RespondWithError(c, utils.BadRequest(err.Error()))
		return
	}
	if err := h.service.MarkAttendance(&attendance); err != nil {
		utils.RespondWithError(c, utils.Internal(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Attendance marked successfully"})

}
