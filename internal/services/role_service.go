package services

import (
	"github.com/maclma/school-management/internal/models"
	"github.com/maclma/school-management/internal/repositories"
)

type RoleService struct {
	repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{repo}
}

func (s *RoleService) CreateRole(role *models.Role) error {
	return s.repo.Create(role)
}

func (s *RoleService) GetRoles() ([]models.Role, error) {
	return s.repo.GetAll()
}

func (s *RoleService) GetRoleByID(id uint) (models.Role, error) {
	return s.repo.GetByID(id)
}

func (s *RoleService) UpdateRole(role *models.Role) error {
	return s.repo.Update(role)
}

func (s *RoleService) DeleteRole(id uint) error {
	return s.repo.Delete(id)
}
