package main

import (
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
