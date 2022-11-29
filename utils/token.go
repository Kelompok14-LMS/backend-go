package utils

import (
	"time"

	"github.com/Kelompok14-LMS/backend-go/pkg"

	"github.com/golang-jwt/jwt"
)

type JWTConfig struct {
	// JWT signature secret
	JWTSecret string
}

func NewJWTConfig(secret string) *JWTConfig {
	return &JWTConfig{
		JWTSecret: secret,
	}
}

type JWTCustomClaims struct {
	UserId   string `json:"user_id"`
	MenteeId string `json:"mentee_id,omitempty"`
	MentorId string `json:"mentor_id,omitempty"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func (config *JWTConfig) GenerateToken(userId string, actorId string, role string) (string, error) {
	claims := JWTCustomClaims{
		UserId: userId,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(1 * time.Hour).Unix(),
		},
	}

	switch role {
	case "mentee":
		claims.MenteeId = actorId
	case "mentor":
		claims.MentorId = actorId
	default:
		return "", pkg.ErrInvalidJWTPayload
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := config.JWTSecret

	return token.SignedString([]byte(jwtSecret))
}
