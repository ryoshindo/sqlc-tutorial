package model

type AccountPassword struct {
	Id        string
	AccountId string
	Hash      []byte
	Salt      []byte
	CreatedAt string
	UpdatedAt string
}
