package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//GetRequestIdFromHeader will read a header and get the request ID
func GetRequestIdFromHeader(req *http.Request, res *echo.Response) string {
	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}
	return id
}
