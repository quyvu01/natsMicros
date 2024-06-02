package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"natsMicros/centralize/api/controllers"
	"natsMicros/centralize/application/configurations"
	"natsMicros/centralize/infrastructure/installers"
)

func main() {
	app := fx.New(
		fx.Provide(installers.NewConfiguration),
		fx.Provide(installers.NewApiService),
		fx.Provide(installers.NewEchoGroup),
		fx.Provide(installers.NewNatsInstaller),
		fx.Provide(controllers.NewMasterDataController),
		fx.Invoke(func(e *echo.Echo, c *configurations.Configuration, nc *nats.Conn, _ *controllers.MasterDataController) {
			go func() {
				e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", c.Port)))
			}()
		}))
	app.Run()
}
