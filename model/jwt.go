package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// JSONWebTokens - represent structure which contains all used in application tokens.
// Such as refresh and access tokens
type JSONWebTokens struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// JWTClaims - default claims for any jwt token
type JWTClaims struct {
	GUID     uuid.UUID
	UserUUID uuid.UUID

	jwt.StandardClaims
}
