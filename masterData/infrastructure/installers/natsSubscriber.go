package installers

import (
	"encoding/json"
	"github.com/ahmetb/go-linq"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"log"
	"natsMicros/buildingBlocks/application/abstractions"
	commonResponse "natsMicros/buildingBlocks/application/responses"
	"natsMicros/contracts/masterData/modelIds"
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
		response, err := json.Marshal(provincesResponse)
		_ = conn.Publish(msg.Reply, response)
	})

	_, _ = conn.Subscribe("masterData.CreateProvinceCommand", func(msg *nats.Msg) {
		province := domain.Province{
			Id: modelIds.NewProvinceId(uuid.New()),
		}
		err := json.Unmarshal(msg.Data, &province)
		if err != nil {
			log.Println("Error unmarshalling createProvince ", err)
			response, _ := json.Marshal(err)
			_ = conn.Publish(msg.Reply, response)
			return
		}
		_, err = repository.CreateOne(&province)
		if err != nil {
			log.Println("Error creating province ", err)
			response, _ := json.Marshal(err)
			_ = conn.Publish(msg.Reply, response)
			return
		}
		_ = conn.Publish(msg.Reply, nil)
	})
	return &NatsSubscriber{conn: conn}
}
