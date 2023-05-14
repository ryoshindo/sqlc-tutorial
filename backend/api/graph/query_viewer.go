package graph

import (
	"context"
	"fmt"

	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph/model"
	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph/session"
)

func (r *queryResolver) Viewer(ctx context.Context) (*model.Account, error) {
	account := session.AccountFromSession(ctx)
	if account == nil {
		return nil, fmt.Errorf("unauthorized")
	}

	return &model.Account{
		ID:    account.ID,
		Email: account.Email,
	}, nil
}
