package services

import (
	"seatmap-backend/entities"
	"time"
)

type UserReader interface {
	Get(username string) (u *entities.User, err error)
	List() (user *[]entities.User,err error)
}

type UserWriter interface {
	Create(userInput *entities.User) (*entities.User, error)
	Delete(username string) (error)
	Update(userInput *entities.User) (error)
}

type UserRepository interface {
	UserReader
	UserWriter
}

type UserService interface {
	GetUser(username string) (u *entities.User, err error)
	ListUsers() (user *[]entities.User,err error)
	CreateUser(userInput *User) (*entities.User, error) 
	VerifyUser(username string, userInput User) (bool, error)
	DeleteUser(username string) (error)
	UpdateUser(userInput *User) (error)
}

type User struct {
	FullName  string
	Username  string
	Password  string
	PasswordConfirmation string
	Salt      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}