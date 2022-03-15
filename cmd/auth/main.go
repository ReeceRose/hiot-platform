package main

import (
	"context"

	"github.com/reecerose/hiot-platform/auth/http"
	zap_middleware "github.com/reecerose/hiot-platform/internal/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var (
	Version        = "dev"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

func main() {
	logger, undo := initLogger()
	defer logger.Sync()
	defer undo()
	zap.L().Info("Starting authentication service",
		zap.String("version", Version),
		zap.String("commit-hash", CommitHash),
		zap.String("build-timestamp", BuildTimestamp),
	)

	e := echo.New()
	e.HideBanner = true
	e.Use(zap_middleware.ZapLogger(zap.L()))
	e.GET("/user/login", http.Login)
	e.GET("/user/register", http.Register)
	e.GET("/token/validate", http.Validate)
	e.GET("/token/issue", http.IssueToken)
	e.Logger.Fatal(e.Start(":80"))
	defer e.Shutdown(context.Background())
}

func initLogger() (*zap.Logger, func()) {
	var logger *zap.Logger
	if Version == "dev" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	undo := zap.ReplaceGlobals(logger)
	return logger, undo
}
