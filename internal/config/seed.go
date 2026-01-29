package config

import (
	"log"

	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedRolesAndAdmin creates default roles and an Admin user
func SeedRolesAndAdmin(db *gorm.DB) {
	roleRepo := repositories.NewRoleRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// default roles
	roles := []string{"Admin", "Teacher", "Student", "Parent"}

	for _, r := range roles {
		existing := db.First(&models.Role{}, "name = ?", r)
		if existing.RowsAffected == 0 {
			roleRepo.Create(&models.Role{Name: r})
			log.Printf("Created role: %s\n", r)
		}
	}

	// default admin user
	var admin models.User
	db.First(&admin, "email = ?", "admin@example.com")
	if admin.ID == 0 {
		password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := models.User{
			FullName: "Admin User",
			Email:    "admin@example.com",
			Password: string(password),
			RoleID:   1, // assuming Admin is first role created
		}
		userRepo.Create(&admin)
		log.Println("Created default Admin user: admin@example.com / admin123")
	}
}
