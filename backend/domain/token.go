package domain

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

const (
	tokenLength = 32
)

const (
	VerifyEmail TokenPurpose = iota + 1
	ResetPassword
	RememberLogin
)

var (
	tokenExpiration = map[TokenPurpose]time.Duration{
		VerifyEmail:   30 * 24 * time.Hour,
		ResetPassword: 3 * time.Hour,
		RememberLogin: 30 * 24 * time.Hour,
	}
)

type TokenPurpose int

type Token struct {
	Id        string       `db:"id"`
	UserId    int          `db:"user_id"`
	Purpose   TokenPurpose `db:"purpose"`
	ExpiresAt time.Time    `db:"expires_at"`
}

func (t Token) IsExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}

func CreateToken(purpose TokenPurpose, userId int) Token {
	return Token{
		Id:        generateTokenId(),
		UserId:    userId,
		Purpose:   purpose,
		ExpiresAt: time.Now().Add(tokenExpiration[purpose]),
	}
}

func generateTokenId() string {
	b := make([]byte, tokenLength/2)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
