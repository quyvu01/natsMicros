package main

import (
	"go.uber.org/fx"
	"natsMicros/buildingBlocks/application/abstractions"
	"natsMicros/buildingBlocks/application/responses"
	"natsMicros/buildingBlocks/infrastructure/repositories"
	"natsMicros/buildingBlocks/infrastructure/services"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	masterDataResponse "natsMicros/contracts/masterData/responses"
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
		fx.Invoke(RunApplication))
	app.Run()
}
