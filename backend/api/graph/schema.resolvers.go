package graph

import "github.com/ryoshindo/sqlc-tutorial/backend/api/graph/generated"

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
