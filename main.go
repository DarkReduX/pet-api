package main

import (
	"context"
	"log/slog"

	"github.com/DarkReduX/pet-api/internal/config"
	"github.com/DarkReduX/pet-api/internal/handler"
	"github.com/DarkReduX/pet-api/internal/repository"
	"github.com/DarkReduX/pet-api/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	slog.Info("App start")

	jwtCfg, err := config.NewJWT()
	if err != nil {
		slog.Error("Failed to initialize configuration: ", slog.String("error", err.Error()))
		return
	}

	ctx := context.Background()

	postgresCfg, err := config.NewPostgres()
	if err != nil {
		slog.Error("Failed to initialize configuration: ", slog.String("error", err.Error()))
		return
	}

	dbPool, err := pgxpool.New(ctx, postgresCfg.URL)
	if err != nil {
		slog.Error("Failed to connect to database: ", slog.String("error", err.Error()))
		return
	}

	userRep := repository.NewUser(dbPool)
	authRep := repository.NewJwtPostgres(dbPool)

	authSvc := service.NewAuth(jwtCfg, userRep, authRep)

	authHandler := handler.NewAuth(authSvc)

	slog.Info("App start")

	svr := echo.New()

	svr.POST("/signup", authHandler.SignUp)
	svr.POST("/signin", authHandler.SignIn)

	svr.Logger.Fatal(svr.Start(":8080"))
}
