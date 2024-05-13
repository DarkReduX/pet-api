package middleware

import (
	"net/http"
	"petProject/model"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const UserIDCtxKey = "user_uuid"

// AuthenticateToken - middleware to validate refresh or access tokens
func AuthenticateToken(config *middleware.JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := jwt.ParseWithClaims(
				strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer "),
				&model.JWTClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return config.SigningKey, nil
				},
			)
			if err != nil {
				log.Errorf("Couldn't parse user auth token: %v", err)
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					err.Error(),
				)
			}

			if !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token expired")
			}

			claims, ok := token.Claims.(*model.JWTClaims)
			if !ok {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Invalid token claims",
				)
			}
			c.Set(UserIDCtxKey, claims.UserUUID)

			return next(c)
		}
	}
}
