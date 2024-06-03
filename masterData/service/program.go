package main

import (
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"natsMicros/buildingBlocks/application/abstractions"
	buildingBlockConfiguration "natsMicros/buildingBlocks/application/configurations"
	"natsMicros/buildingBlocks/infrastructure/repositories"
	"natsMicros/masterData/application/configurations"
	"natsMicros/masterData/domain"
	"natsMicros/masterData/infrastructure/installers"
)

func main() {
	app := fx.New(
		fx.Provide(installers.NewConfiguration),
		fx.Provide(installers.NewNatsInstaller),
		fx.Provide(installers.NewNatsSubscriber),
		fx.Provide(installers.NewMongoDbSettingConfig),
		fx.Provide(installers.NewMongoDbClient),
		fx.Provide(installers.NewMongoDbConnectionName[domain.Province]),
		fx.Provide(fx.Annotate(repositories.NewMongoDbRepository[domain.Province], fx.As(new(abstractions.IDatabaseRepository[domain.Province])))),
		fx.Invoke(func(_ *configurations.Configuration, _ *nats.Conn, _ *installers.NatsSubscriber, _ *mongo.Client, _ *buildingBlockConfiguration.MongoDbCollectionSetting[domain.Province]) {

		}))
	app.Run()
}
