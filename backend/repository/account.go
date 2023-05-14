package repository

import (
	"context"

	"github.com/ryoshindo/sqlc-tutorial/backend/model"
)

type AccountRepository interface {
	FindById(ctx context.Context, id string) (*model.Account, error)
	FindByEmail(ctx context.Context, email string) (*model.Account, error)
	FindPasswordsByAccountId(ctx context.Context, accountId string) ([]model.AccountPassword, error)
}
