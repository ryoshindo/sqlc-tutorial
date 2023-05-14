package graph

import "github.com/ryoshindo/sqlc-tutorial/backend/repository"

type Resolver struct {
	accountRepo repository.AccountRepository
}

func NewResolver(accountRepo repository.AccountRepository) *Resolver {
	return &Resolver{
		accountRepo: accountRepo,
	}
}
