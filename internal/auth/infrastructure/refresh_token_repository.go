// internal/auth/infrastructure/refresh_token_repository.go
package infrastructure

import (
	"gin-quickstart/internal/auth/domain"

	"github.com/jmoiron/sqlx"
)

type refreshTokenRepository struct {
    db *sqlx.DB
}

func NewRefreshTokenRepository(db *sqlx.DB) domain.RefreshTokenRepository {
    return &refreshTokenRepository{db: db}
}

func (r *refreshTokenRepository) Create(token *domain.RefreshToken) error {
    query := `INSERT INTO refresh_tokens (user_id, token_hash, expires_at) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(query, token.UserID, token.TokenHash, token.ExpiresAt)
    return err
}

func (r *refreshTokenRepository) GetByHash(hash string) (*domain.RefreshToken, error) {
    var token domain.RefreshToken
    query := `SELECT * FROM refresh_tokens WHERE token_hash = $1 AND is_revoked = false AND expires_at > NOW()`
    err := r.db.Get(&token, query, hash)
    return &token, err
}

func (r *refreshTokenRepository) RevokeByHash(hash string) error {
    query := `UPDATE refresh_tokens SET is_revoked = true WHERE token_hash = $1`
    _, err := r.db.Exec(query, hash)
    return err
}

func (r *refreshTokenRepository) DeleteExpired() error {
    query := `DELETE FROM refresh_tokens WHERE expires_at < NOW()`
    _, err := r.db.Exec(query)
    return err
}

