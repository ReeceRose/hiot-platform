package main

import (
	"github.com/labstack/echo/v4"
	"github.com/reecerose/hiot-platform/auth/controller"
	"github.com/reecerose/hiot-platform/internal/http/server"
	"github.com/reecerose/hiot-platform/internal/logger"
	"github.com/reecerose/hiot-platform/internal/utils"
	"github.com/reecerose/hiot-platform/pkg/types"
	"github.com/sirupsen/logrus"
)

var (
	Version        = "dev"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

const (
	fallbackJWTSecret    = "u8x/A?D(G+KbPeShVmYq3t6v9y$B&E)H"
	fallbackPort         = "443"
	fallbackCertDir      = "certs"
	fallbackCrtFile      = "localhost.crt"
	fallbackKeyFile      = "localhost.key"
	environmentJWTSecret = "AUTH_JWT_SECRET"
	environmentPort      = "AUTH_HTTPS_PORT"
	environmentCertsDir  = "CERTS_DIR"
	environmentCrtFile   = "CRT_FILE"
	environmentKeyFile   = "KEY_FILE"
)

type config struct {
	jwtSecret string
	port      string
	certDir   string
	crtFile   string
	keyFile   string
}

func main() {
	config := initConfig()
	logger := logger.Instance()

	logger.WithFields(logrus.Fields{
		"version":         Version,
		"commit-hash":     CommitHash,
		"build-timestamp": BuildTimestamp,
	}).Info("Starting authentication service")

	httpServer := server.New()
	httpServer.Setup(&server.ServerConfig{
		Port:    config.port,
		CertDir: config.certDir,
		CrtFile: config.crtFile,
		KeyFile: config.keyFile,
	}, func() {
		setupHTTPRoutes(httpServer.Instance(), config)
	})
	defer httpServer.Shutdown()
}

func initConfig() *config {
	return &config{
		jwtSecret: utils.GetEnvironmentVariable(environmentJWTSecret, fallbackJWTSecret),
		port:      utils.GetEnvironmentVariable(environmentPort, fallbackPort),
		certDir:   utils.GetEnvironmentVariable(environmentCertsDir, fallbackCertDir),
		crtFile:   utils.GetEnvironmentVariable(environmentCrtFile, fallbackCrtFile),
		keyFile:   utils.GetEnvironmentVariable(environmentKeyFile, fallbackKeyFile),
	}
}

func setupHTTPRoutes(e *echo.Echo, config *config) {
	controller := controller.NewAuthController(&types.AuthConfig{
		SecretKey: config.jwtSecret,
	})
	e.GET("/token/validate", func(c echo.Context) error { return controller.Validate(c) })
	e.GET("/token/issue", func(c echo.Context) error { return controller.Issue(c) })
}
