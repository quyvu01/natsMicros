package installers

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/thoas/go-funk"
	"log"
	"natsMicros/buildingBlocks/application/abstractions"
	commonResponse "natsMicros/buildingBlocks/application/responses"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	"natsMicros/contracts/masterData/responses"
	"natsMicros/masterData/domain"
)

type NatsSubscriber struct {
	conn *nats.Conn
}

func NewNatsSubscriber(conn *nats.Conn, repository abstractions.IDatabaseRepository[domain.Province]) *NatsSubscriber {
	_, _ = conn.Subscribe("masterData.GetProvincesQuery", func(msg *nats.Msg) {
		getProvinceQuery := getProvinces.GetProvincesQuery{}
		err := json.Unmarshal(msg.Data, &getProvinceQuery)
		if err != nil {
			log.Println("Error unmarshalling getProvinces ", err)
			response, _ := json.Marshal(commonResponse.PaginationResponse[responses.ProvinceResponse]{})
			_ = conn.Publish(msg.Reply, response)
			return
		}
		provinces, err := repository.GetManyByCondition(nil)
		if err != nil {
			log.Println("Error getting provinces: ", err)
			response, _ := json.Marshal(commonResponse.PaginationResponse[responses.ProvinceResponse]{})
			_ = conn.Publish(msg.Reply, response)
			return
		}

		provincesResponse := commonResponse.PaginationResponse[responses.ProvinceResponse]{
			Items: funk.Map(provinces, func(province domain.Province) responses.ProvinceResponse {
				return responses.ProvinceResponse{Name: province.Name, ModelResponse: commonResponse.ModelResponse{Id: province.Id}}
			}).([]responses.ProvinceResponse),
			CurrentPageIndex: 1,
			TotalPage:        1,
			TotalRecord:      1,
		}
		response, err := json.Marshal(provincesResponse)
		_ = conn.Publish(msg.Reply, response)
	})
	return &NatsSubscriber{conn: conn}
}
