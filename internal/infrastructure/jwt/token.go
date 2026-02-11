package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWTクレーム
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken JWTトークンを生成
func GenerateToken(userID uint, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", jwt.ErrSignatureInvalid
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken JWTトークンを検証
func ValidateToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, jwt.ErrSignatureInvalid
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
