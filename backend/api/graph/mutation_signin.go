package graph

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph/model"
	m "github.com/ryoshindo/sqlc-tutorial/backend/model"
)

func (r *mutationResolver) Signin(ctx context.Context, input model.SigninInput) (bool, error) {
	account, err := r.accountRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return false, err
	}

	if account == nil {
		return false, nil
	}

	passwords, err := r.accountRepo.FindPasswordsByAccountID(ctx, account.ID)
	if err != nil {
		return false, err
	}

	for _, password := range passwords {
		input, err := m.RawPassword(input.Password).CalculateHash(password.Salt)
		if err != nil {
			return false, err
		}

		if bytes.Equal(input, password.Hash) {
			return true, nil
		}
	}

	return false, fmt.Errorf("failed password authentication")
}
