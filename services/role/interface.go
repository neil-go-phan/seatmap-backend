package services

import (
	"seatmap-backend/entities"
)

type RoleReader interface {
	Get(roleName string) (r *entities.Role, err error)
	List() (role *[]entities.Role,err error)
}

type RoleWriter interface {
	// Create(userInput *entities.User) (*entities.User, error)
}

type RoleRepository interface {
	RoleReader
	RoleWriter
}

type RoleService interface {
	GetRole(roleName string) (r *entities.Role, err error)
	ListRole() (role *[]entities.Role, err error) 
	// CreateUser(userInput *User) (*entities.User, error) 
	// VerifyUser(username string, userInput User) (bool, error)
	// UpdateUser(e *entity.User) error
	// DeleteUser(id entity.ID) error
}

type Role struct {
	RoleName string
}