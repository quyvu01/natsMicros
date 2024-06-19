package main

import (
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"natsMicros/buildingBlocks/application/abstractions"
	buildingBlockConfiguration "natsMicros/buildingBlocks/application/configurations"
	"natsMicros/buildingBlocks/application/responses"
	"natsMicros/buildingBlocks/infrastructure/repositories"
	"natsMicros/buildingBlocks/infrastructure/services"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	masterDataResponse "natsMicros/contracts/masterData/responses"
	"natsMicros/masterData/application/configurations"
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
			fx.Annotate(repositories.NewMongoDbRepository[domain.Province], fx.As(new(abstractions.IDatabaseRepository[domain.Province]))),
			fx.Annotate(services.NewNatResponseService[getProvinces.GetProvincesQuery, responses.PaginationResponse[masterDataResponse.ProvinceResponse]],
				fx.As(new(abstractions.IMessageResponse[getProvinces.GetProvincesQuery, responses.PaginationResponse[masterDataResponse.ProvinceResponse]])))),
		fx.Invoke(func(_ *configurations.Configuration, _ *nats.Conn, _ *installers.NatsSubscriber, _ *mongo.Client, _ *buildingBlockConfiguration.MongoDbCollectionSetting[domain.Province]) {
		}))
	app.Run()
}
