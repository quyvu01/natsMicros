package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"natsMicros/centralize/api/controllers"
	"natsMicros/centralize/application/configurations"
	"time"
)

const (
	MaxHeaderBytes = 1 << 20
	ReadTimeout    = 15 * time.Second
	WriteTimeout   = 15 * time.Second
)

func RunApplication(lc fx.Lifecycle, echo *echo.Echo, c *configurations.Configuration, _ *nats.Conn, _ *controllers.MasterDataController) {
	echo.Server.ReadTimeout = ReadTimeout
	echo.Server.WriteTimeout = WriteTimeout
	echo.Server.MaxHeaderBytes = MaxHeaderBytes
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				echo.Logger.Fatal(echo.Start(fmt.Sprintf(":%s", c.Port)))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return echo.Shutdown(ctx)
		},
	})
}
