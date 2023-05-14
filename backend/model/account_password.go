package model

import (
	"time"

	"golang.org/x/crypto/scrypt"
)

type AccountPassword struct {
	ID        string
	AccountID string
	Hash      []byte
	Salt      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RawPassword string

func (p RawPassword) CalculateHash(salt []byte) ([]byte, error) {
	key, err := scrypt.Key([]byte(p), salt, 1<<15, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return key, nil
}
