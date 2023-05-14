package graph

import (
	"context"
	"fmt"

	"github.com/ryoshindo/sqlc-tutorial/backend/api/graph/model"
)

func (r *queryResolver) Viewer(ctx context.Context) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}
