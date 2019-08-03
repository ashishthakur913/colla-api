package apphtml

import (
	"github.com/labstack/echo"
	"net/http"
)

func HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home", "<p>HTML Test</p>")
}