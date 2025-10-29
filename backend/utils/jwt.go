package utils

import (
	"fmt"
	"time"

	
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTClams represents the data stored in the JWT
type JWTClaims struct {
	UserID 	 uuid.UUID `json:"user_id"`
	Email  	 string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// this generates a JWT for a given user ID and email
func GenerateJWT(userID uuid.UUID, email, username, secret string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore:  jwt.NewNumericDate(time.Now()),
		},
	}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			return "", fmt.Errorf("failed to sign token: %w", err)
		}
	return tokenString, nil
}

// ValidateJWT validates the JWT and returns the claims if valid
func ValidateJWT(tokenString, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		 // verify signing method
		 if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		 }
		 return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")

}