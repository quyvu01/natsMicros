package installers

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"natsMicros/centralize/application/configurations"
)

func NewNatsInstaller(lc fx.Lifecycle, configuration *configurations.Configuration) *nats.Conn {
	nc, err := nats.Connect(configuration.NatsSetting.Url)
	if err != nil {
		panic(fmt.Errorf("unable to connect to NatS: %v", err))
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if nc != nil {
				nc.Close()
			}
			return nil
		},
	})
	return nc
}
