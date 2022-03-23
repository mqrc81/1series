package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHttpError(errorType HttpErrorType, error error, details ...interface{}) *echo.HTTPError {
	httpError := HttpError{
		Type:    errorType.string(),
		Error:   error.Error(),
		Details: http.StatusText(errorType.statusCode()),
	}
	if len(details) > 0 {
		httpError.Details = details[0]
	}
	return echo.NewHTTPError(errorType.statusCode(), httpError)
}

type HttpError struct {
	Type    string
	Error   string
	Details interface{}
}

type HttpErrorType interface {
	string() string
	statusCode() int
}

// HttpApplicationErrorType is an error aimed at the user.
//  - can't be fixed by a developer
//  - can be fixed by a user.
type HttpApplicationErrorType int

const (
	_ HttpApplicationErrorType = iota
	Parameter
	Form
	Authentication
	Database
	Unknown
)

var httpApplicationErrorTypes = [...]string{"Z-Parameter", "Z-Form", "Z-Authentication", "Z-Database", "Z-Unknown"}

func (et HttpApplicationErrorType) string() string {
	return httpApplicationErrorTypes[et-1]
}

func (et HttpApplicationErrorType) statusCode() int {
	return http.StatusBadRequest
}

// HttpSystemErrorType is an error aimed at an external system.
//  - can't be fixed by a developer.
//  - can't be fixed by a user.
type HttpSystemErrorType int

const (
	_ HttpSystemErrorType = iota
	Tmdb
	Trakt
)

var httpSystemErrorTypes = [...]string{"Z-Tmdb", "Z-Trakt"}

func (et HttpSystemErrorType) string() string {
	return httpSystemErrorTypes[et-1]
}

func (et HttpSystemErrorType) statusCode() int {
	return http.StatusInternalServerError
}

// HttpConsistencyErrorType is an error aimed at the developer.
//  - can be fixed by a developer.
//  - can't be fixed by a user.
type HttpConsistencyErrorType int

const (
	_ HttpConsistencyErrorType = iota
	Unsupported
	Unimplemented
)

var httpConsistencyErrorTypes = [...]string{"Z-Unsupported", "Z-Unimplemented"}

func (et HttpConsistencyErrorType) string() string {
	return httpConsistencyErrorTypes[et-1]
}

func (et HttpConsistencyErrorType) statusCode() int {
	return http.StatusInternalServerError
}
