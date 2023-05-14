package repository

import (
	"context"

	"github.com/ryoshindo/sqlc-tutorial/backend/model"
)

type AccountRepository interface {
	FindByID(ctx context.Context, id string) (*model.Account, error)
	FindByEmail(ctx context.Context, email string) (*model.Account, error)
	FindPasswordsByAccountID(ctx context.Context, accountId string) ([]model.AccountPassword, error)
	Create(ctx context.Context, account *model.Account) error
}
