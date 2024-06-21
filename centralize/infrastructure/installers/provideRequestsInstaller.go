package installers

import (
	"go.uber.org/fx"
	"natsMicros/buildingBlocks/application/abstractions"
	"natsMicros/buildingBlocks/application/responses"
	"natsMicros/buildingBlocks/infrastructure/services"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	masterDataResponse "natsMicros/contracts/masterData/responses"
)

func AnnotateRequestService[TRequest any, TResponse any]() interface{} {
	return fx.Annotate(services.NewNatRequestService[TRequest, TResponse], fx.As(new(abstractions.IMessageRequest[TRequest, TResponse])))
}

func ProvideRequestsServices() fx.Option {
	var options []interface{}
	getProvinceOption := AnnotateRequestService[getProvinces.GetProvincesQuery, responses.PaginationResponse[masterDataResponse.ProvinceResponse]]()
	options = append(options, getProvinceOption)
	return fx.Provide(options...)
}
