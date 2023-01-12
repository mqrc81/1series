package domain

import "time"

type Token struct {
	Token     string       `db:"token"`
	UserId    int          `db:"user_id"`
	Purpose   TokenPurpose `db:"purpose"`
	ExpiresAt time.Time    `db:"expires_at"`
}

type TokenPurpose int

const (
	VerifyEmail TokenPurpose = iota + 1
	ResetPassword
)
