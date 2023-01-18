package errors

import (
	"database/sql"
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stoewer/go-strcase"
	"net/http"
	"strings"
)

type ApiError = *echo.HTTPError

type Params = map[string]interface{}

type clientError struct {
	Message string `json:"errorMessage,omitempty"`
}

func Internal(err error) ApiError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

func Unauthorized() ApiError {
	return echo.NewHTTPError(http.StatusUnauthorized, clientError{Message: "Session has expired. Please log back in."})
}

func NotFound(entity string, parameters Params) ApiError {
	var parameterDetails string
	if len(parameters) > 0 {
		var parameterPairs []string
		for key, val := range parameters {
			parameterPairs = append(parameterPairs, fmt.Sprintf("%v '%v'", key, val))
		}
		parameterDetails = " with " + strings.Join(parameterPairs, ", ")
	}
	// e.g. "with id 13, name Vikings"
	return echo.NewHTTPError(http.StatusNotFound, clientError{Message: fmt.Sprintf("Could not find %v%v.", entity, parameterDetails)})
}

func MissingParameter(parameter string) ApiError {
	return echo.NewHTTPError(http.StatusBadRequest, clientError{Message: fmt.Sprintf("Missing %v parameter.", parameter)})
}

func InvalidParam(message string) ApiError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, clientError{Message: message})
}

func AdminOnly() ApiError {
	return echo.NewHTTPError(http.StatusForbidden, clientError{Message: "Only the big boss is allowed to perform this action."})
}

// Conditional api error types depending on error content

func FromDatabase(err error, entity string, parameters Params) ApiError {
	if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		return NotFound(entity, parameters)
	}
	return Internal(err)
}

func FromTmdb(err error, entity string, parameters Params) ApiError {
	if tmdbErr, ok := err.(tmdb.Error); ok && tmdbErr.StatusCode == 34 {
		return NotFound(entity, parameters)
	}
	return Internal(err)
}

func FromValidation(err error) ApiError {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var parameters []string
		for _, validationError := range validationErrors {
			parameters = append(parameters, strcase.KebabCase(validationError.StructField()))
		}
		return InvalidParam(fmt.Sprintf("Invalid parameters %v.", strings.Join(parameters, ", ")))
	}
	return InvalidParam("Invalid parameters.")
}

func FromEmail(err error) ApiError {
	// TODO ms
	return Internal(err)
}
