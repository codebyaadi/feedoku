package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey []byte

func init() {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	jwtKey = []byte(jwtSecretKey)
}

type Claims struct {
	UserID uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	ApiKey string    `json:"api_key"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId uuid.UUID, name, email, apiKey string) (string, string, error) {
	// Access Token ( short lived )
	accessExpirationTime := time.Now().Add(15 * time.Minute)
	accessClaims := &Claims{
		UserID: userId,
		Name:   name,
		Email:  email,
		ApiKey: apiKey,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: accessExpirationTime},
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token ( long lived )
	refreshExpirationTime := time.Now().Add(30 * 24 * time.Hour)
	refreshClaims := &Claims{
		UserID: userId,
		Name:   name,
		Email:  email,
		ApiKey: apiKey,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: refreshExpirationTime},
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	err = SetRefreshToken(userId, signedRefreshToken, time.Until(refreshExpirationTime))
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
