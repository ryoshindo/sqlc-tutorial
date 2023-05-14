package sqlc

import (
	"context"

	"github.com/ryoshindo/sqlc-tutorial/backend/model"
)

type SessionRepository struct{}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{}
}

func (r *SessionRepository) FindByToken(ctx context.Context, token string) (*model.Session, error) {
	panic("not implemented")
}

func (r *SessionRepository) Create(ctx context.Context, session *model.Session) error {
	panic("not implemented")
}

func (r *SessionRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented")
}
