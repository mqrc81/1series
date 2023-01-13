package domain

import "time"

type Token struct {
	TokenId   string       `db:"token_id"`
	UserId    int          `db:"user_id"`
	Purpose   TokenPurpose `db:"purpose"`
	ExpiresAt time.Time    `db:"expires_at"`
}

type TokenPurpose int

const (
	VerifyEmail TokenPurpose = iota + 1
	ResetPassword
)
