package installers

import (
	"go.uber.org/fx"
	"natsMicros/buildingBlocks/application/abstractions"
	"natsMicros/buildingBlocks/application/responses"
	"natsMicros/buildingBlocks/infrastructure/services"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	masterDataResponse "natsMicros/contracts/masterData/responses"
)

func AnnotatedResponseService[TRequest any, TResponse any]() interface{} {
	return fx.Annotate(services.NewNatResponseService[TRequest, TResponse], fx.As(new(abstractions.IMessageResponse[TRequest, TResponse])))
}

func ProvideResponsesServices() fx.Option {
	var options []interface{}
	getProvinceOption := AnnotatedResponseService[getProvinces.GetProvincesQuery, responses.PaginationResponse[masterDataResponse.ProvinceResponse]]()
	options = append(options, getProvinceOption)
	return fx.Provide(options...)
}
