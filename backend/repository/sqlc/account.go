package sqlc

import (
	"context"

	"github.com/ryoshindo/sqlc-tutorial/backend/model"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) FindByID(ctx context.Context, id string) (*model.Account, error) {
	panic("not implemented")
}

func (r *AccountRepository) FindByEmail(ctx context.Context, email string) (*model.Account, error) {
	panic("not implemented")
}

func (r *AccountRepository) FindPasswordsByAccountID(ctx context.Context, accountId string) ([]model.AccountPassword, error) {
	panic("not implemented")
}

func (r *AccountRepository) Create(ctx context.Context, account *model.Account) error {
	panic("not implemented")
}
