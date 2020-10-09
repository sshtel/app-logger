package handler_v1
import (
	"net/http"
	"github.com/labstack/echo"
)

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
