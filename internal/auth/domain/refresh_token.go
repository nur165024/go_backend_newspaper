package domain

import "time"

type RefreshToken struct {
    ID        int       `db:"id" json:"id"`
    UserID    int       `db:"user_id" json:"user_id"`
    TokenHash string    `db:"token_hash" json:"token_hash"`
    ExpiresAt time.Time `db:"expires_at" json:"expires_at"`
    IsRevoked bool      `db:"is_revoked" json:"is_revoked"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type RefreshTokenRepository interface {
    Create(token *RefreshToken) error
    GetByHash(hash string) (*RefreshToken, error)
    RevokeByHash(hash string) error
    DeleteExpired() error
}
