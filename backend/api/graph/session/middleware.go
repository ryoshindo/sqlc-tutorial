package session

import (
	"context"
	"net/http"
	"os"

	"github.com/ryoshindo/sqlc-tutorial/backend/model"
	"github.com/ryoshindo/sqlc-tutorial/backend/repository"
)

var (
	accountCtxKey        = &contextKey{"account"}
	sessionCtxKey        = &contextKey{"session"}
	sessionHandlerCtxKey = &contextKey{"session_handler"}
)

type contextKey struct {
	name string
}

func Middleware(
	accountRepo repository.AccountRepository,
	sessionRepo repository.SessionRepository,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			srw := &sessionResponseWriter{w, nil}
			sessionHandler := &sessionHandler{
				accountRepo,
				sessionRepo,
				srw,
			}

			ctx := context.WithValue(r.Context(), sessionHandlerCtxKey, sessionHandler)
			c, err := r.Cookie(getCookieKey())

			if err != nil || c == nil {
				next.ServeHTTP(srw, r.WithContext(ctx))
				return
			}

			session, err := sessionRepo.FindByToken(ctx, c.Value)
			if err != nil {
				next.ServeHTTP(srw, r.WithContext(ctx))
				return
			}

			account, err := accountRepo.FindByID(ctx, session.AccountID)
			if err != nil {
				next.ServeHTTP(srw, r.WithContext(ctx))
				return
			}

			ctx = context.WithValue(ctx, accountCtxKey, account)
			ctx = context.WithValue(ctx, sessionCtxKey, session)

			next.ServeHTTP(srw, r.WithContext(ctx))
		})
	}
}

type sessionHandler struct {
	accountRepo repository.AccountRepository
	sessionRepo repository.SessionRepository

	responseWriter *sessionResponseWriter
}

type sessionResponseWriter struct {
	http.ResponseWriter

	session *model.Session
}

func getCookieKey() string {
	if value, ok := os.LookupEnv("SESSION_TOKEN_COOKIE_KEY"); ok {
		return value
	}
	return "sqlc_tutorial_session_token"
}

func (w *sessionResponseWriter) Write(b []byte) (int, error) {
	if w.session != nil {
		http.SetCookie(w, &http.Cookie{
			Name:     getCookieKey(),
			Value:    w.session.Token,
			HttpOnly: true,
		})
	}

	return w.ResponseWriter.Write(b)
}

func AccountFromSession(ctx context.Context) *model.Account {
	account, _ := ctx.Value(accountCtxKey).(*model.Account)
	return account
}

func CreateSession(ctx context.Context, account *model.Account) error {
	handler, _ := ctx.Value(sessionHandlerCtxKey).(*sessionHandler)
	if handler == nil {
		return nil
	}

	return handler.CreateSession(ctx, account)
}

func DeleteSession(ctx context.Context) error {
	handler, _ := ctx.Value(sessionHandlerCtxKey).(*sessionHandler)
	if handler == nil {
		return nil
	}
	session, _ := ctx.Value(sessionCtxKey).(*model.Session)
	if session == nil {
		return nil
	}

	return handler.DeleteSession(ctx, session)
}

func (h *sessionHandler) CreateSession(ctx context.Context, account *model.Account) error {
	session, err := model.NewSession(account.ID)
	if err != nil {
		return err
	}

	if err := h.sessionRepo.Create(ctx, session); err != nil {
		return err
	}
	h.responseWriter.session = session

	return nil
}

func (h *sessionHandler) DeleteSession(ctx context.Context, session *model.Session) error {
	if err := h.sessionRepo.Delete(ctx, session.ID); err != nil {
		return err
	}

	return nil
}
