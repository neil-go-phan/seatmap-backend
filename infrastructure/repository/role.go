package repository

import (
	"seatmap-backend/entities"

	"gorm.io/gorm"
)

type RoleRepo struct {
	DB *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		DB: db,
	}
}

func (repo *RoleRepo)Get(roleName string) (r *entities.Role, err error) {
	role := new(entities.Role)
	err = repo.DB.Where(map[string]interface{}{"role_name": roleName}).Find(&role).Error
	if err!= nil {
		return nil, err
	}
	return role, nil
}

func (repo *RoleRepo)List() (role *[]entities.Role,err error) {
	roles := make([]entities.Role, 10)
	err = repo.DB.Select("role_name").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return &roles, nil
}