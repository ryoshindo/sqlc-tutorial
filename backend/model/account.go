package model

type Account struct {
	Id        string
	Email     string
	CreatedAt string
	UpdatedAt string
}

type AccountPassword struct {
	Id        string
	AccountId string
	Hash      []byte
	Salt      []byte
	CreatedAt string
	UpdatedAt string
}
