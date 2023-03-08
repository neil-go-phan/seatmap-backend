package services

import "seatmap-backend/entities"

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