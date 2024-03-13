package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
