package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func CreateServer() *echo.Echo {
	e := echo.New()
	e.GET("/hello", hello)
	return e
}

func hello(c echo.Context) error {
	message := fmt.Sprintf("Hello %s.", c.QueryParam("name"))
	return c.String(http.StatusOK, message)
}
