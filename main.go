package main

import (
	"context"
	"log/slog"
	"petProject/internal/config"
	"petProject/internal/handler"
	"petProject/internal/repository"
	"petProject/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	jwtCfg, err := config.NewJWT()
	if err != nil {
		slog.Error("Failed to initialize configuration: ", slog.String("error", err.Error()))
	}

	ctx := context.Background()
	slog.Info("App start")

	postgresCfg, err := config.NewPostgreSQLEnv()
	if err != nil {
		slog.Error("Failed to initialize configuration: ", slog.String("error", err.Error()))
	}

	dbPool, err := pgxpool.New(ctx, postgresCfg.URL)
	if err != nil {
		slog.Error("Failed to connect to database: ", slog.String("error", err.Error()))
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
