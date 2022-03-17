package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/reecerose/hiot-platform/auth/service"
	"github.com/reecerose/hiot-platform/internal/logger"
	"github.com/reecerose/hiot-platform/internal/utils"
	"github.com/reecerose/hiot-platform/pkg/types"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	config  *types.AuthConfig
	service types.IAuthService
	log     *logrus.Logger
}

//NewAuthController returns a new initialized auth controller
func NewAuthController(config *types.AuthConfig) *AuthController {
	service := service.NewAuthService(config)
	return &AuthController{
		config:  config,
		service: service,
		log:     logger.Instance(),
	}
}

//Validate a JWT or API token
func (controller *AuthController) Validate(c echo.Context) error {
	requestId := utils.GetRequestIdFromHeader(c.Request(), c.Response())
	tokenString := c.Request().Header.Get("Authorization")
	controller.log.WithFields(logrus.Fields{
		"token": tokenString,
		"id":    requestId,
	}).Info("Attempting to validate token")

	if tokenString == "" {
		message := "No token set via Authorization header"
		controller.log.WithFields(logrus.Fields{
			"token": tokenString,
			"id":    requestId,
		}).Info(message)
		return c.JSON(400, types.StandardResponse{
			StatusCode: 400,
			Error:      message,
			Success:    false,
			Data:       nil,
		})
	}

	response := controller.service.Validate(tokenString)
	message := "succesfully validated token"
	if len(response.Error) != 0 {
		message = response.Error
	}
	controller.log.WithFields(logrus.Fields{
		"token": tokenString,
		"id":    requestId,
	}).Info(message)
	return c.JSON(response.StatusCode, response)
}

//Issue a JWT or API token
func (controller *AuthController) Issue(c echo.Context) error {
	return nil
}
