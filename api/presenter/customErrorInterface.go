package presenter

import "net/http"

type ErrorReponse struct {
	ErrLog             error
	ErrorConvention
}

// ErrorConvention is used to declare ErrorName with 
type ErrorConvention struct {
	ErrorName string
	ErrorMessage	string	`json:"message"`
	HttpStatusCode uint16
}

func (e ErrorReponse)Error() string {
	return e.ErrorName
}

var ERROR_VALIDATE_TOKEN_FAIL = ErrorConvention{
	ErrorName: "Validate tolen fail",
	ErrorMessage: "Unauthorized access",
	HttpStatusCode: http.StatusUnauthorized,
}

var ERROR_NO_REFESH_TOKEN = ErrorConvention{
	ErrorName: "No refresh token",
	ErrorMessage: "No refresh token provided",
	HttpStatusCode: http.StatusUnauthorized,
}

var ERROR_GENERATE_TOKEN_FAIL = ErrorConvention{
	ErrorName: "Cant generate token",
	ErrorMessage: "Fail to generate token",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_NO_PERMISSION = ErrorConvention{
	ErrorName: "No permisstion",
	ErrorMessage: "No permission granted",
	HttpStatusCode: http.StatusBadRequest,
}

var ERROR_BAD_REQUEST = ErrorConvention{
	ErrorName: "Bad request",
	ErrorMessage: "Bad request",
	HttpStatusCode: http.StatusBadRequest,
}

var ERROR_INPUT_INVALID = ErrorConvention{
	ErrorName: "Input invalid",
	ErrorMessage: "Input invalid",
	HttpStatusCode: http.StatusBadRequest,
}

var ERROR_SIGNIN_INCORRECT = ErrorConvention{
	ErrorName: "Sign in incorrect",
	ErrorMessage: "Username or password is incorrect",
	HttpStatusCode: http.StatusBadRequest,
}

var ERROR_USERNAME_TAKEN = ErrorConvention{
	ErrorName: "Username is already taken",
	ErrorMessage: "Username is already taken",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_UPDATE_FAIL = ErrorConvention{
	ErrorName: "Update fail",
	ErrorMessage: "Fail to update",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_DELETE_FAIL = ErrorConvention{
	ErrorName: "Delete fail",
	ErrorMessage: "Fail to delete",
	HttpStatusCode: http.StatusInternalServerError,
}

var ERROR_SERVER_ERROR = ErrorConvention{
	ErrorName: "Server error",
	ErrorMessage: "Server error",
	HttpStatusCode: http.StatusInternalServerError,
}