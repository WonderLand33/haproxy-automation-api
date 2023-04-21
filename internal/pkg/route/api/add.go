package api

import (
	"haproxy-automation-api/internal/pkg/haproxy"
	"haproxy-automation-api/internal/pkg/server"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AddRequest struct {
	IP string
}

func Add(c echo.Context) error {
	req := new(AddRequest)
	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ip := strings.TrimSpace(req.IP)

	if ip == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "IP为空")
	}

	ips, err := haproxy.GetBlockedIPs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, server.Response{
			Msg: err.Error(),
		})
	}

	if ips.Has(ip) {
		return c.JSON(http.StatusOK, server.Response{
			Msg: "已经有该IP，无需添加",
		})
	}

	ips.Add(ip)

	err = haproxy.Update(ips)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, server.Response{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, server.Response{
		Msg: "it works",
	})
}
