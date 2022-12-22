package usecases

type (
	RegisterForm struct {
		Username string `json:"username" validate:"required,alphanum,min=3,max=16"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=3"`
	}

	LoginForm struct {
		EmailOrUsername string `json:"email-or-username" validate:"required,email|alphanum"`
		Password        string `json:"password" validate:"required"`
	}
)
