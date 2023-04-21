package api

import (
	"github.com/labstack/echo/v4"
	"haproxy-automation-api/internal/pkg/haproxy"
	"haproxy-automation-api/internal/pkg/server"
	"net/http"
)

func Reload(c echo.Context) error {
	err := haproxy.Reload()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, server.Response{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, server.Response{
		Msg: "reload ok",
	})
}
