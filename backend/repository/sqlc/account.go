package sqlc

import (
	"context"
	"time"

	"github.com/ryoshindo/sqlc-tutorial/backend/db"
	"github.com/ryoshindo/sqlc-tutorial/backend/model"
	generatedsqlc "github.com/ryoshindo/sqlc-tutorial/backend/model/sqlc"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) FindByID(ctx context.Context, id string) (*model.Account, error) {
	sAccount, err := db.Get(ctx).FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	account := &model.Account{
		ID:        sAccount.ID,
		Email:     sAccount.Email,
		CreatedAt: sAccount.CreatedAt,
		UpdatedAt: sAccount.UpdatedAt,
	}

	return account, nil
}

func (r *AccountRepository) FindByEmail(ctx context.Context, email string) (*model.Account, error) {
	sAccount, err := db.Get(ctx).FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	account := &model.Account{
		ID:        sAccount.ID,
		Email:     sAccount.Email,
		CreatedAt: sAccount.CreatedAt,
		UpdatedAt: sAccount.UpdatedAt,
	}

	return account, nil
}

func (r *AccountRepository) FindPasswordsByAccountID(ctx context.Context, accountId string) ([]model.AccountPassword, error) {
	sPasswords, err := db.Get(ctx).FindPasswordsByAccountId(ctx, accountId)
	if err != nil {
		return nil, err
	}

	passwords := make([]model.AccountPassword, len(sPasswords))
	for i, sPassword := range sPasswords {
		passwords[i] = model.AccountPassword{
			ID:        sPassword.ID,
			AccountID: sPassword.AccountID,
			Hash:      sPassword.Hash,
			Salt:      sPassword.Salt,
			CreatedAt: sPassword.CreatedAt,
			UpdatedAt: sPassword.UpdatedAt,
		}
	}

	return passwords, nil
}

func (r *AccountRepository) Create(ctx context.Context, account *model.Account) error {
	now := time.Now()
	params := generatedsqlc.CreateAccountParams{
		ID:        account.ID,
		Email:     account.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := db.Get(ctx).CreateAccount(ctx, params); err != nil {
		return err
	}

	return nil
}
