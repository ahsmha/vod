package auth

import (
	"errors"
	"fmt"
	"time"

	"vod/internal/config"

	"github.com/dgrijalva/jwt-go"
)

// AuthService represents the authentication and authorization service.
type AuthService struct {
	config *config.Config
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService(config *config.Config) *AuthService {
	return &AuthService{
		config: config,
	}
}

// GenerateJWT generates a new JSON Web Token (JWT) for the given user ID.
func (a *AuthService) GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Duration(a.config.JWTExpirationMinutes) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(a.config.JWTSecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %v", err)
	}

	return signedToken, nil
}

// VerifyJWT verifies the authenticity of a JSON Web Token (JWT).
func (a *AuthService) VerifyJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}

		// Return the secret key used for signing
		return []byte(a.config.JWTSecretKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to verify JWT: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract and return the user ID from the token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", errors.New("invalid user ID in token claims")
		}
		return userID, nil
	}

	return "", errors.New("failed to extract user ID from token claims")
}
