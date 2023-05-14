package graph

import "github.com/ryoshindo/sqlc-tutorial/backend/repository"

type Resolver struct {
	accountRepo repository.AccountRepository
	sessionRepo repository.SessionRepository
}

func NewResolver(accountRepo repository.AccountRepository, sessionRepo repository.SessionRepository) *Resolver {
	return &Resolver{
		accountRepo: accountRepo,
		sessionRepo: sessionRepo,
	}
}
