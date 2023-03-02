package handler

import (
	"errors"
	"regexp"
	"seatmap-backend/api/presenter"
	"seatmap-backend/entities"
	"time"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

// func (u User) TableName() string {
// 	return "users"
// }
// type ID = uuid.UUID

// //NewID create a new entity ID
// func NewID() ID {
// 	return ID(uuid.New())
// }

func ToUserEntities() {

}

func NewUser(userInput *presenter.User) (*entities.User, error) {
	user := &entities.User{
		FullName:  userInput.FullName,
		Username:  userInput.Username,
		Password:  userInput.Password,
		Role:      userInput.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := Validate(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Validate(user *entities.User) error {
	err := ValidateFullname(user)
	if err != nil {
		return err
	}
	err = ValidateUsernameAndPassword(user)
	if err != nil {
		return err
	}
	return nil
}

func ValidateFullname(user *entities.User) error {
	validate = validator.New()
	err := validate.Var(user.FullName, "required,min=8,max=50")
	if err != nil {
		return err
	}
	checkRegexFullName := checkRegexp(user.FullName, "full_name")
	if !checkRegexFullName {
		return errors.New("full name must not contain special character")
	}
	return nil
}

func ValidateUsernameAndPassword(user *entities.User) error {
	validate = validator.New()
	match := checkRegexp(user.Password, "usernameAndPassword")
	if !match {
		return errors.New("password must not contain special character")
	}
	match = checkRegexp(user.Username, "usernameAndPassword")
	if !match {
		return errors.New("username must not contain special character")
	}
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	return nil
}

func checkRegexp(checkedString string, checkType string) bool {
	switch checkType {
	case "usernameAndPassword":
		match, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", checkedString)
		return match
	case "full_name":
		match, _ := regexp.MatchString("^[a-zA-Z0-9_ ]*$", checkedString)
		return match
	}
	return false
}
