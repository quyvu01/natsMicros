package main

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"natsMicros/masterData/application/configurations"
	"natsMicros/masterData/infrastructure/installers"
)

func main() {
	app := fx.New(
		fx.Provide(installers.NewConfiguration),
		fx.Provide(installers.NewNatsInstaller),
		fx.Provide(installers.NewNatsSubscriber),
		fx.Invoke(func(_ *configurations.Configuration, _ *nats.Conn, _ *installers.NatsSubscriber) {

		}))
	app.Run()
}
