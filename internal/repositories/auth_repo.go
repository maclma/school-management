package repositories

import (
	"github.com/maclma/school-management/internal/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// Find user by email
func (r *AuthRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find user by ID
func (r *AuthRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").
		First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
