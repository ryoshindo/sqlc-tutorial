package model

type Account struct {
	ID        string
	Email     string
	CreatedAt string
	UpdatedAt string
}

func NewAccount() *Account {
	return &Account{
		ID: NewUlidString(),
	}
}
