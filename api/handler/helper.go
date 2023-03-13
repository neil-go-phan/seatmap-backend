package handler

import (
	"crypto/rand"
	"errors"
	"regexp"
	"seatmap-backend/api/presenter"
	"seatmap-backend/services"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
)

var validate *validator.Validate
var TOKEN_SERECT_KEY = []byte("GolenOwl2023")

// TODO: create a env.example contain all const
const ACCESS_TOKEN_LIFE = 5 * time.Minute // 5 min
const REFRESH_TOKEN_LIFE = 24 * time.Hour // 1 day
const RANDOM_TOKEN_STRING_SIZE = 8

func newServicesUser(userInput *presenter.User) *services.User {
	user := &services.User{
		FullName:             userInput.FullName,
		Username:             userInput.Username,
		Password:             userInput.Password,
		PasswordConfirmation: userInput.PasswordConfirmation,
		Role:                 "Staff",
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	return user
}
// just return errorName, error handler middleware will handle return
func NewErrorReponse(err error, errorName string) *presenter.ErrorReponse {
	errorConvention := generateErrorProperty(errorName)
	e := &presenter.ErrorReponse{
		ErrLog:         err,
		ErrorConvention: errorConvention,
	}
	return e
}

func generateErrorProperty(errorName string) (presenter.ErrorConvention) {
	switch errorName {
	case presenter.ERROR_VALIDATE_TOKEN_FAIL.ErrorName:
		return presenter.ERROR_VALIDATE_TOKEN_FAIL

	case presenter.ERROR_NO_REFESH_TOKEN.ErrorName:
		return presenter.ERROR_NO_REFESH_TOKEN

	case presenter.ERROR_GENERATE_TOKEN_FAIL.ErrorName:
		return presenter.ERROR_GENERATE_TOKEN_FAIL

	case presenter.ERROR_NO_PERMISSION.ErrorName:
		return presenter.ERROR_NO_PERMISSION

	case presenter.ERROR_BAD_REQUEST.ErrorName:
		return presenter.ERROR_BAD_REQUEST

	case presenter.ERROR_INPUT_INVALID.ErrorName:
		return presenter.ERROR_INPUT_INVALID

	case presenter.ERROR_SIGNIN_INCORRECT.ErrorName:
		return presenter.ERROR_SIGNIN_INCORRECT

	}
	return presenter.ERROR_SERVER_ERROR;
}

func validateSignUp(user *services.User) error {
	err := validateFullname(user)
	if err != nil {
		return err
	}
	err = validateUsernameAndPassword(user)
	if err != nil {
		return err
	}
	return nil
}

func validateFullname(user *services.User) error {
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

func validateUsernameAndPassword(user *services.User) error {
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

func generateAccessToken(username, role string) (string, error) {
	expirationTime := time.Now().Add(ACCESS_TOKEN_LIFE)

	claims := &presenter.JWTClaim{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(TOKEN_SERECT_KEY)
}

func GenerateRefreshToken(username, role string) (string, error) {
	randomString, err := generateRandomTokenString()
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(REFRESH_TOKEN_LIFE)
	claims := &presenter.JWTClaim{
		Username:     username,
		Role:         role,
		RandomString: randomString,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(TOKEN_SERECT_KEY)
}

func generateRandomTokenString() ([]byte, error) {
	var randomString = make([]byte, RANDOM_TOKEN_STRING_SIZE)

	_, err := rand.Read(randomString[:])

	if err != nil {
		return nil, err
	}

	return randomString, nil
}
