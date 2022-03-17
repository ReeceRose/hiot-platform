package logger

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/reecerose/hiot-platform/internal/utils"
	"github.com/sirupsen/logrus"
)

//LoggerMiddleware is a middleware that logs HTTP requests.
func LoggerMiddleware(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			var err error
			var errorMessage string
			if err = next(c); err != nil {
				c.Error(err)
				errorMessage = err.Error()
			}
			stop := time.Now()

			id := utils.GetRequestIdFromHeader(c.Request(), c.Response())
			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}
			fields := logrus.Fields{
				"id":         id,
				"remote-ip":  c.RealIP(),
				"host":       req.Host,
				"method":     req.Method,
				"uri":        req.RequestURI,
				"status":     res.Status,
				"error":      errorMessage,
				"size":       reqSize,
				"latency":    stop.Sub(start).String(),
				"referer":    req.Referer(),
				"user-agent": req.UserAgent(),
			}

			n := res.Status
			switch {
			case n >= 500:
				logger.WithFields(fields).Error("")
			case n >= 400:
				logger.WithFields(fields).Warn("")
			case n >= 300:
				logger.WithFields(fields).Info("")
			default:
				logger.WithFields(fields).Info("")
			}

			return err
		}
	}
}
