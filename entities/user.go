package models

import (
	"errors"
	"regexp"
	"time"
	// "github.com/google/uuid"
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username" validate:"required,min=8,max=16"`
	Password  string    `json:"password" validate:"required,md5"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var validate *validator.Validate

func (u User) TableName() string {
	return "users"
}
// type ID = uuid.UUID

// //NewID create a new entity ID
// func NewID() ID {
// 	return ID(uuid.New())
// }

// func NewUser(fullname, username, password, role string) (*User, error) {
// 	u := &User{
// 		ID:        NewID(),
// 		FullName:     fullname,
// 		Username: username,
// 		Password:  password,
// 		Role: role,
// 		CreatedAt: time.Now(),
// 		UpdatedAt : time.Now(),
// 	}
// 	pwd, err := generatePassword(password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	u.Password = pwd
// 	err = u.Validate()
// 	if err != nil {
// 		return nil, ErrInvalidEntity
// 	}
// 	return u, nil
// }

func (user *User) ValidateFullname() error {
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

func (user *User) ValidateUsernameAndPassword() error {
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
