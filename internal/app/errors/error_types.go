package errors

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	LogMsg  string `json:"-"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewCustomError(code int, message string, logMsg string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		LogMsg:  logMsg,
	}
}

func ConflictError(message string) *CustomError {
	return NewCustomError(409, message, "Conflict error occurred")
}

func AuthError(message string) *CustomError {
	return NewCustomError(401, message, "Authentication failed")
}

func BadRequestError(message string) *CustomError {
	return NewCustomError(400, message, "Bad request")
}

func NotFoundError(message string) *CustomError {
	return NewCustomError(404, message, "Not found")
}

func InternalServerError(err error) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		LogMsg:  fmt.Sprintf("Detailed error: %v", err),
	}
}
