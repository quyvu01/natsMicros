package installers

import (
	"github.com/ahmetb/go-linq"
	"log"
	"natsMicros/buildingBlocks/application/abstractions"
	commonResponse "natsMicros/buildingBlocks/application/responses"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	"natsMicros/contracts/masterData/responses"
	"natsMicros/masterData/domain"
)

type NatsSubscriber struct {
}

func NewNatsSubscriber(repository abstractions.IDatabaseRepository[domain.Province],
	getProvinceResponse abstractions.IMessageResponse[getProvinces.GetProvincesQuery, commonResponse.PaginationResponse[responses.ProvinceResponse]]) *NatsSubscriber {
	_ = getProvinceResponse.Response(func(getProvincesQuery getProvinces.GetProvincesQuery) (commonResponse.PaginationResponse[responses.ProvinceResponse], error) {
		provinces, err := repository.GetManyByCondition(nil)
		if err != nil {
			log.Println("Error getting provinces: ", err)
			return commonResponse.PaginationResponse[responses.ProvinceResponse]{}, err
		}
		pResponse := make([]responses.ProvinceResponse, len(provinces))
		linq.From(provinces).SelectT(func(i *domain.Province) responses.ProvinceResponse {
			return responses.ProvinceResponse{Name: i.Name, ModelResponse: commonResponse.ModelResponse{Id: i.Id.String()}}
		}).ToSlice(&pResponse)
		provincesResponse := commonResponse.PaginationResponse[responses.ProvinceResponse]{
			Items:            pResponse,
			CurrentPageIndex: 1,
			TotalPage:        1,
			TotalRecord:      1,
		}
		return provincesResponse, nil
	})
	return &NatsSubscriber{}
}
