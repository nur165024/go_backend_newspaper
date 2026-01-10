// pkg/auth/jwt.go - Updated version
package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTSecret struct {
	secretKey string
	accessTokenExpireMinutes int
	refreshTokenExpireDays int
}

type Claims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	TokenType string `json:"token_type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func NewJWTServices(secretKey string, accessMinutes, refreshDays int) *JWTSecret {
	return &JWTSecret{
		secretKey: secretKey,
		accessTokenExpireMinutes: accessMinutes,
		refreshTokenExpireDays: refreshDays,
	}
}

// Generate token pair with env-based expiration
func (j *JWTSecret) GenerateTokenPair(id int, name, email, userName string) (*TokenPair, error) {
    // Only generate JWT tokens
    accessToken, err := j.generateToken(id, name, email, userName, "access", 
        time.Now().Add(time.Duration(j.accessTokenExpireMinutes)*time.Minute))
    if err != nil {
        return nil, err
    }

    refreshToken, err := j.generateToken(id, name, email, userName, "refresh", 
        time.Now().Add(time.Duration(j.refreshTokenExpireDays)*24*time.Hour))
    if err != nil {
        return nil, err
    }

    return &TokenPair{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        ExpiresIn:    int64(j.accessTokenExpireMinutes * 60),
    }, nil
}


func (j *JWTSecret) generateToken(id int, name, email, userName, tokenType string, expireTime time.Time) (string, error) {
	claims := &Claims{
		ID:        id,
		Name:      name,
		Email:     email,
		UserName:  userName,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTSecret) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// Generate refresh token hash for database storage
func GenerateRefreshTokenHash() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:]), nil
}
