package model

import "time"

type Account struct {
	ID        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount() *Account {
	return &Account{
		ID: NewUlidString(),
	}
}
