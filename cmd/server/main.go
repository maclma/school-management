package main

import (
	"log"

	"github.com/maclma/school-management/internal/config"
	"github.com/maclma/school-management/internal/db"
	"github.com/maclma/school-management/internal/handlers"
	"github.com/maclma/school-management/internal/middleware"
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
	"github.com/maclma/school-management/internal/routes"
	"github.com/maclma/school-management/internal/seed"
	"github.com/maclma/school-management/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("config error", err)
	}

	// Set Gin mode
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to database
	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatal("db error", err)
	}

	// Run migrations
	database.AutoMigrate(
		&models.Student{},
		&models.Teacher{},
		&models.Class{},
		&models.Subject{},
		&models.Attendance{},
		&models.User{},
		&models.Role{},
	)

	// Seed roles and admin user
	config.SeedRolesAndAdmin(database)
	if err := seed.Run(database); err != nil {
		log.Fatal("seeding error", err)
	}

	// ===== Repositories =====
	studentRepo := repositories.NewStudentRepository(database)
	teacherRepo := repositories.NewTeacherRepository(database)
	classRepo := repositories.NewClassRepository(database)
	subjectRepo := repositories.NewSubjectRepository(database)
	attendanceRepo := repositories.NewAttendanceRepository(database)
	roleRepo := repositories.NewRoleRepository(database)
	userRepo := repositories.NewUserRepository(database)

	// ===== Services =====
	studentService := services.NewStudentService(studentRepo)
	teacherService := services.NewTeacherService(teacherRepo)
	classService := services.NewClassService(classRepo)
	subjectService := services.NewSubjectService(subjectRepo)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	roleService := services.NewRoleService(roleRepo)
	userService := services.NewUserService(userRepo)

	// Auth service
	authService := services.NewAuthService(userService, cfg.JWTSecret)

	// ===== Handlers =====
	studentHandler := handlers.NewStudentHandler(studentService)
	teacherHandler := handlers.NewTeacherHandler(teacherService)
	classHandler := handlers.NewClassHandler(classService)
	subjectHandler := handlers.NewSubjectHandler(subjectService)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)
	roleHandler := handlers.NewRoleHandler(roleService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	// ===== Router =====
	router := gin.Default()
	router.SetTrustedProxies(nil) // avoid proxy warning
	router.Use(middleware.CORSMiddleware())

	// Set JWT secret in middleware
	middleware.SetJWTSecret(cfg.JWTSecret)

	// Register routes
	routes.RegisterRoutes(
		router,
		studentHandler,
		teacherHandler,
		classHandler,
		subjectHandler,
		attendanceHandler,
		userHandler,
		roleHandler,
		authHandler,
		cfg.JWTSecret,
	)

	// Start server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port " + port)
	router.Run(":" + port)
}