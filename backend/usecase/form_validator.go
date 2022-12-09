package usecase

import (
	"errors"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

type Form interface {
	Validate() error
}

type formErrors map[string]string

type RegisterForm struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`

	UsernameTaken bool
	EmailTaken    bool
}

func (f *RegisterForm) Validate() error {
	formErrors := formErrors{}

	// TODO validate

	if len(formErrors) == 0 {
		return nil
	}
	errorMsg := "Registration-form is invalid\n"
	for _, formError := range formErrors {
		errorMsg += formError + "\n"
	}
	return errors.New(errorMsg)
}

type LoginForm struct {
	// Identifier can be either username or email
	Identifier string `form:"identifier"`
	Password   string `form:"password"`

	IdentifierNotFound bool
}

func (f *LoginForm) Validate() error {
	formErrors := formErrors{}

	// TODO validate

	if len(formErrors) == 0 {
		return nil
	}
	errorMsg := "Login-form is invalid\n"
	for _, formError := range formErrors {
		errorMsg += formError + "\n"
	}
	return errors.New(errorMsg)
}
