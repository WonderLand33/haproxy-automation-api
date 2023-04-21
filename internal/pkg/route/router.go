package route

import (
	"haproxy-automation-api/internal/pkg/route/api"

	"github.com/labstack/echo/v4"
)

func Install(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/banned-ips", api.List)
	v1.POST("/banned-ip", api.Add)
	v1.DELETE("/banned-ip", api.Del)
	v1.POST("/reload", api.Reload)
}
