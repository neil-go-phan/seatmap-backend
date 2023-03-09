package services

import (
	"errors"
	"seatmap-backend/entities"
)

type roleService struct {
	repo RoleRepository
}


func NewRoleService(r RoleRepository) *roleService {
	return &roleService{
		repo: r,
	}
}

func (s *roleService) GetRole(roleName string) (*entities.Role, error) {
	return s.repo.Get(roleName)
}

func (s *roleService) ListRole() (role *[]entities.Role, err error) {
	return s.repo.List()
}

func (s *roleService) Validate(roleName string) (err error) {
	role, err := s.repo.Get(roleName)
	if err != nil {
		return err
	}
	if role.RoleName == "" {
		return errors.New("Role invalid")
	}
	return nil
}