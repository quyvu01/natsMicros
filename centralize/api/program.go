package main

import (
	"go.uber.org/fx"
	"natsMicros/buildingBlocks/application/abstractions"
	commonResponse "natsMicros/buildingBlocks/application/responses"
	"natsMicros/buildingBlocks/infrastructure/services"
	"natsMicros/centralize/api/controllers"
	"natsMicros/centralize/infrastructure/configurations"
	"natsMicros/centralize/infrastructure/installers"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	"natsMicros/contracts/masterData/responses"
)

func main() {
	app := fx.New(
		fx.Provide(installers.NewConfiguration,
			installers.NewEchoServer,
			installers.NewEchoGroup,
			installers.NewNatsInstaller,
			controllers.NewMasterDataController, fx.Annotate(services.NewNatRequestService[getProvinces.GetProvincesQuery, commonResponse.PaginationResponse[responses.ProvinceResponse]],
				fx.As(new(abstractions.IMessageRequest[getProvinces.GetProvincesQuery, commonResponse.PaginationResponse[responses.ProvinceResponse]])))),
		fx.Invoke(RunApplication, configurations.MiddlewareConfiguration))
	app.Run()
}
