package handler

import (
	"log/slog"
	"net/http"

	"github.com/DarkReduX/pet-api/internal/service"
	"github.com/DarkReduX/pet-api/model"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	svc *service.Auth
}

func NewAuth(svc *service.Auth) *Auth {
	return &Auth{svc: svc}
}

func (h *Auth) SignUp(c echo.Context) error {
	ctx := c.Request().Context()

	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		slog.Error("Could not bind user: ", slog.String("error", err.Error()))
		return c.NoContent(http.StatusBadRequest)
	}

	user, tokens, err := h.svc.SignUp(ctx, user)
	if err != nil {
		slog.Error("Could not sign up user: ", slog.String("error", err.Error()))
		return c.NoContent(http.StatusBadRequest)
	}

	response := model.AuthResponse{
		User:   user,
		Tokens: tokens,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Auth) SignIn(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(model.SignInRequest)
	if err := c.Bind(&req); err != nil {
		slog.Error("Could not bind user: ", slog.String("error", err.Error()))
		return c.NoContent(http.StatusBadRequest)
	}

	user, tokens, err := h.svc.SignIn(ctx, req)
	if err != nil {
		slog.Error("Could not sign in user: ", slog.String("error", err.Error()))
		return c.NoContent(http.StatusBadRequest)
	}

	response := model.AuthResponse{
		User:   user,
		Tokens: tokens,
	}

	return c.JSON(http.StatusOK, response)
}
