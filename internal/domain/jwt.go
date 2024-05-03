package domain

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ExtractUserIDFromToken(tokenString string) (uint64, error)
}

type JWTTokenService struct {
	secretKey []byte
}

func NewJWTTokenService(secretKey string) *JWTTokenService {
	return &JWTTokenService{
		secretKey: []byte(secretKey),
	}
}

func (s *JWTTokenService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *JWTTokenService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return token, nil
}

func (s *JWTTokenService) ExtractUserIDFromToken(tokenString string) (uint, error) {
	token, err := s.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}
	return uint(userID), nil
}
