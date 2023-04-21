package api

import (
	"haproxy-automation-api/internal/pkg/haproxy"
	"haproxy-automation-api/internal/pkg/server"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type DelRequest struct {
	IP string
}

func Del(c echo.Context) error {
	req := new(DelRequest)
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

	if !ips.Has(ip) {
		return c.JSON(http.StatusOK, server.Response{
			Msg: "没有该IP，无需删除",
		})
	}

	ips.Remove(ip)

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
