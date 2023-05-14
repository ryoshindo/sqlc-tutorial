package sqlc

import (
	"context"

	"github.com/ryoshindo/sqlc-tutorial/backend/model"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) FindById(ctx context.Context, id string) (*model.Account, error) {
	panic("not implemented")
}

func (r *AccountRepository) FindByEmail(ctx context.Context, email string) (*model.Account, error) {
	panic("not implemented")
}

func (r *AccountRepository) FindPasswordsByAccountId(ctx context.Context, accountId string) ([]model.AccountPassword, error) {
	panic("not implemented")
}
