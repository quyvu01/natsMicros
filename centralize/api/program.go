package main

import (
	"go.uber.org/fx"
	"natsMicros/centralize/api/controllers"
	"natsMicros/centralize/infrastructure/configurations"
	"natsMicros/centralize/infrastructure/installers"
)

func main() {
	app := fx.New(
		fx.Provide(installers.NewConfiguration,
			installers.NewEchoServer,
			installers.NewEchoGroup,
			installers.NewNatsInstaller,
			controllers.NewMasterDataController),
		installers.ProvideRequestsServices(),
		fx.Invoke(RunApplication, configurations.MiddlewareConfiguration, configurations.SwaggerConfiguration))
	app.Run()
}
