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
		provinces, err := repository.GetPaginationByCondition(nil)
		if err != nil {
			log.Println("Error getting provinces: ", err)
			return commonResponse.PaginationResponse[responses.ProvinceResponse]{}, err
		}
		pResponse := make([]responses.ProvinceResponse, len(provinces.Items))
		linq.From(provinces.Items).SelectT(func(i *domain.Province) responses.ProvinceResponse {
			return responses.ProvinceResponse{Name: i.Name, ModelResponse: commonResponse.ModelResponse{Id: i.Id.String()}}
		}).ToSlice(&pResponse)
		provincesResponse := commonResponse.PaginationResponse[responses.ProvinceResponse]{
			Items:            pResponse,
			CurrentPageIndex: provinces.CurrentPageIndex,
			TotalPage:        provinces.TotalPage,
			TotalRecord:      provinces.TotalRecord,
		}
		return provincesResponse, nil
	})
	return &NatsSubscriber{}
}
