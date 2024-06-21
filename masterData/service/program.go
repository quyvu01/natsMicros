package main

import (
	"go.uber.org/fx"
	"natsMicros/buildingBlocks/application/abstractions"
	"natsMicros/buildingBlocks/infrastructure/repositories"
	"natsMicros/masterData/domain"
	"natsMicros/masterData/infrastructure/installers"
)

func main() {
	app := fx.New(
		fx.Provide(installers.NewConfiguration,
			installers.NewNatsInstaller,
			installers.NewNatsSubscriber,
			installers.NewMongoDbSettingConfig,
			installers.NewMongoDbClient,
			installers.NewMongoDbConnectionName[domain.Province],
			fx.Annotate(repositories.NewMongoDbRepository[domain.Province], fx.As(new(abstractions.IDatabaseRepository[domain.Province])))),
		installers.ProvideResponsesServices(),
		fx.Invoke(RunApplication))
	app.Run()
}
