package user

import "seatmap-backend/entities"

type Reader interface {
	Get(username string) (u *entities.User, err error)
	List() (user *[]entities.User,err error)
}

type Writer interface {
	Create(userInput *entities.User) (*entities.User, error)
}

type UserRepository interface {
	Reader
	Writer
}

type UserUsecase interface {
	GetUser(username string) (u *entities.User, err error)
	ListUsers() (user *[]entities.User,err error)
	CreateUser(userInput *entities.User) (*entities.User, error)
	// UpdateUser(e *entity.User) error
	// DeleteUser(id entity.ID) error
}