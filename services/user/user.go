package services

import (
	"encoding/base64"
	"errors"
	"seatmap-backend/entities"
)

var SALT_SIZE uint8 = 8 // 8 byte

type userService struct {
	repo UserRepository
}


func NewService(r UserRepository) *userService {
	return &userService{
		repo: r,
	}
}

func (s *userService) GetUser(username string) (*entities.User, error) {
	return s.repo.Get(username)
}

func (s *userService) ListUsers() (user *[]entities.User,err error) {
	return s.repo.List()
}

func (s *userService) CreateUser(userInput *User) (*entities.User, error) {
	if userInput.Password != userInput.PasswordConfirmation {
		return nil, errors.New("password confirm is not match")
	}
	salt, err := generateRandomSalt(SALT_SIZE)
	if err != nil {
		return nil, errors.New("error when generate salt")
	}
// hash password with salt and argon2
	hashedPassword,err := hashPassword(userInput.Password, salt)
	if err != nil {
		return nil, err
	}
	userInput.Password = hashedPassword
	userInput.Salt = base64.RawStdEncoding.EncodeToString(salt)
	entitiesUser := NewEntitiesUser(userInput)
	return s.repo.Create(entitiesUser)
}

func (s *userService) VerifyUser(username string, userInput User) (bool, error) {
	if username != userInput.Username {
		return false, errors.New("username is incorrect")
	}
	userFromDB, err := s.GetUser(username)
	if err != nil {
		return false, err
	}
	return verifyPassword(userInput.Password, userFromDB.Password)
}