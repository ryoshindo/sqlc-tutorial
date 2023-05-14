package model

import (
	"time"

	"github.com/ryoshindo/sqlc-tutorial/backend/util/rand"
)

type Session struct {
	ID        string
	AccountID string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	SessionExpiresIn = time.Hour * 24 * 7
)

func NewSession(accountId string) (*Session, error) {
	maxStringLength := 128
	token, err := rand.GenerateRandomString(maxStringLength)
	if err != nil {
		return nil, err
	}

	session := &Session{
		ID:        NewUlidString(),
		AccountID: accountId,
		Token:     token,
		ExpiresAt: time.Now().Add(SessionExpiresIn),
	}

	return session, nil
}
