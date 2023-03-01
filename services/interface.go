package services
import (
	
)
type Reader interface {
	Get(username string) (u *entities.User, err error)
	List() (user *[]entities.User,err error)
}

type Writer interface {
	Create(userInput *entities.User) (*entities.User, error)
}