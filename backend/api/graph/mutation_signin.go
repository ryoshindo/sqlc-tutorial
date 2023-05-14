package graph

import (
	"context"
	"fmt"

	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph/model"
)

func (r *mutationResolver) Signin(ctx context.Context, input model.SigninInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
