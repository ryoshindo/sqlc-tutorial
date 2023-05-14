package sqlc

import (
	"context"
	"time"

	"github.com/ryoshindo/sqlc-tutorial/backend/db"
	"github.com/ryoshindo/sqlc-tutorial/backend/model"
	generatedsqlc "github.com/ryoshindo/sqlc-tutorial/backend/model/sqlc"
)

type SessionRepository struct{}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{}
}

func (r *SessionRepository) FindByToken(ctx context.Context, token string) (*model.Session, error) {
	sSession, err := db.Get(ctx).FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	session := &model.Session{
		ID:        sSession.ID,
		AccountID: sSession.AccountID,
		Token:     sSession.Token,
		ExpiresAt: sSession.ExpiresAt,
		CreatedAt: sSession.CreatedAt,
		UpdatedAt: sSession.UpdatedAt,
	}

	return session, nil
}

func (r *SessionRepository) Create(ctx context.Context, session *model.Session) error {
	now := time.Now()

	params := generatedsqlc.CreateSessionParams{
		ID:        session.ID,
		AccountID: session.AccountID,
		Token:     session.Token,
		ExpiresAt: session.ExpiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := db.Get(ctx).CreateSession(ctx, params); err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) Delete(ctx context.Context, id string) error {
	if err := db.Get(ctx).DeleteSession(ctx, id); err != nil {
		return err
	}

	return nil
}
