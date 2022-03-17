package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reecerose/hiot-platform/internal/logger"
)

// Server is an interface which provides method signatures for an HTTP server
type Server interface {
	Instance() *echo.Echo
	Setup(*ServerConfig, func())
	Shutdown()
}

type ServerConfig struct {
	Port    string
	CertDir string
	CrtFile string
	KeyFile string
}

type echoServer struct {
	Server *echo.Echo
}

var (
	_ Server = (*echoServer)(nil)
)

//New returns a new instance of an echo HTTP server
func New() Server {
	return &echoServer{
		Server: echo.New(),
	}
}

//Instance returns a singleton echo instance
func (s *echoServer) Instance() *echo.Echo {
	return s.Server
}

//Setup the web server
func (s *echoServer) Setup(config *ServerConfig, setupRoutes func()) {
	e := s.Instance()
	e.HideBanner = true
	e.Pre(middleware.HTTPSRedirect())

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(logger.LoggerMiddleware(logger.Instance()))

	setupRoutes()

	e.Logger.Fatal(e.StartTLS(":"+config.Port,
		config.CertDir+"/"+config.CrtFile,
		config.CertDir+"/"+config.KeyFile))
}

//Shutdown the web server
func (s *echoServer) Shutdown() {
	s.Instance().Shutdown(context.Background())
}
