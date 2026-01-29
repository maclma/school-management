package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maclma/school-management/internal/handlers"
	"github.com/maclma/school-management/internal/middleware"
)

func RegisterRoutes(
	r *gin.Engine,
	studentHandler *handlers.StudentHandler,
	teacherHandler *handlers.TeacherHandler,
	classHandler *handlers.ClassHandler,
	subjectHandler *handlers.SubjectHandler,
	attendanceHandler *handlers.AttendanceHandler,
	userHandler *handlers.UserHandler,
	roleHandler *handlers.RoleHandler,
	authHandler *handlers.AuthHandler,
	jwtSecret string,
) {

	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.GET("/me", middleware.AuthMiddleware(), authHandler.Me)
	}

	// Users (protected)
	users := r.Group("/users", middleware.AuthMiddleware())
	{
		users.POST("/", userHandler.Create)
		users.GET("/", userHandler.GetAll)
		users.GET("/:id", userHandler.GetByID)
		users.PUT("/:id", userHandler.Update)
	}

	// Students
	students := r.Group("/students", middleware.AuthMiddleware())
	{
		students.POST("/", studentHandler.CreateStudent)
		students.GET("/", studentHandler.GetAll)
		students.GET("/:id", studentHandler.GetStudentByID)
		students.PUT("/:id", studentHandler.UpdateStudent)
		students.DELETE("/:id", studentHandler.DeleteStudent)
	}

	// Teachers
	teachers := r.Group("/teachers", middleware.AuthMiddleware())
	{
		teachers.POST("/", teacherHandler.CreateTeacher)
		teachers.GET("/", teacherHandler.GetTeachers)
	}

	// Classes
	classes := r.Group("/classes", middleware.AuthMiddleware())
	{
		classes.POST("/", classHandler.CreateClass)
		classes.GET("/", classHandler.GetClasses)
	}

	// Subjects
	subjects := r.Group("/subjects", middleware.AuthMiddleware())
	{
		subjects.POST("/", subjectHandler.CreateSubject)
		subjects.GET("/", subjectHandler.GetSubjects)

	}

	// Attendance
	attendance := r.Group("/attendance", middleware.AuthMiddleware())
	{
		attendance.POST("/", attendanceHandler.CreateAttendance)
		attendance.GET("/", attendanceHandler.GetAllAttendance)
	}
}
