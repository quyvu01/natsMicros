package installers

import (
	"github.com/labstack/echo/v4"
)

func NewEchoServer() *echo.Echo {
	return echo.New()
}
