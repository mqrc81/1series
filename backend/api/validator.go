package api

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

type FormErrors map[string]string

type RegisterForm struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`

	UsernameTaken bool
	EmailTaken    bool
	FormErrors
}

func (f *RegisterForm) Validate() bool {
	f.FormErrors = FormErrors{}

	return len(f.FormErrors) == 0
}

type LoginForm struct {
	// Identifier can be either username or email
	Identifier string `form:"identifier"`
	Password   string `form:"password"`

	IdentifierNotFound bool
	FormErrors
}

func (f *LoginForm) Validate() bool {
	f.FormErrors = FormErrors{}

	return len(f.FormErrors) == 0
}
