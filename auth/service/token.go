package service

import (
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/reecerose/hiot-platform/internal/logger"
	"github.com/reecerose/hiot-platform/pkg/types"
	"github.com/sirupsen/logrus"
)

type authService struct {
	config *types.AuthConfig
	log    *logrus.Logger
}

var (
	_ types.IAuthService = (*authService)(nil)
)

// NewAuthService returns an initialized auth service
func NewAuthService(config *types.AuthConfig) types.IAuthService {
	return &authService{
		config: config,
		log:    logger.Instance(),
	}
}

//Validate a JWT or API token
func (s *authService) Validate(tokenString string) types.StandardResponse {
	if strings.Contains(tokenString, "Bearer") {
		// validate JWT
		success, err := s.validateJWT(tokenString)
		if success {
			return types.StandardResponse{
				StatusCode: 200,
				Data:       nil,
				Success:    true,
			}
		} else {
			return types.StandardResponse{
				StatusCode: 400,
				Data:       nil,
				Error:      err.Error(),
				Success:    true,
			}
		}
	} else if strings.Contains(tokenString, "API") {
		// validate API token
		// TODO: validate by querying database
		return types.StandardResponse{
			StatusCode: 200,
			Data:       nil,
			Success:    true,
		}
	}
	return types.StandardResponse{}
}

//validateJWT validates a given JWT
func (s *authService) validateJWT(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(strings.Replace(tokenString, "Bearer ", "", 1), &types.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SecretKey), nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*types.AuthClaims); ok && token.Valid {
		return true, nil
	}
	return false, nil
}

//Issue a JWT or API token
func (s *authService) Issue() types.StandardResponse {
	return types.StandardResponse{}
}
