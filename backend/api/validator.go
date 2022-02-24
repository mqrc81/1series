package api

import (
	"fmt"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

func NewValidator(store domain.Store) echo.Validator {
	return validator{store}
}

func (v validator) Validate(i interface{}) error {
	return i.(formValidator).validate(v.store)
}

type RegisterForm struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (r RegisterForm) validate(store domain.Store) error {
	errorsAndFields := make(map[string]string)

	if matched, _ := regexp.MatchString(emailRegex, r.Email); !matched {
		errorsAndFields["email"] = "Invalid email"
	}

	return validationErrorFrom(errorsAndFields)
}

type LoginForm struct {
	UsernameOrEmail string `form:"usernameOrEmail"`
	Password        string `form:"password"`
}

func (r LoginForm) validate(store domain.Store) error {
	errorsAndFields := make(map[string]string)

	if r.UsernameOrEmail == "" {
		errorsAndFields["usernameOrEmail"] = "Username or email can not be empty"
	}

	return validationErrorFrom(errorsAndFields)
}

type validator struct {
	store domain.Store
}

type formValidator interface {
	validate(store domain.Store) error
}

type validationError struct {
	errorsAndFields map[string]string
}

func (e validationError) Error() string {
	return fmt.Sprintf("error validating form: %v", e.errorsAndFields)
}

func validationErrorFrom(errorsAndFields map[string]string) error {
	if len(errorsAndFields) > 0 {
		return validationError{errorsAndFields}
	}
	return nil
}
