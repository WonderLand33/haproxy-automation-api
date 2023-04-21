package api

import (
	"haproxy-automation-api/internal/pkg/haproxy"
	"haproxy-automation-api/internal/pkg/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	res, err := haproxy.GetBlockedIPs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, server.Response{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, server.Response{
		Data: map[string]interface{}{
			"ips": res.List(),
		},
		Msg: "it works",
	})
}
