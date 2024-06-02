package installers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"log"
	commonResponse "natsMicros/buildingBlocks/application/responses"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	"natsMicros/contracts/masterData/responses"
)

type NatsSubscriber struct {
	conn *nats.Conn
}

func NewNatsSubscriber(conn *nats.Conn) *NatsSubscriber {
	_, _ = conn.Subscribe("masterData.GetProvincesQuery", func(msg *nats.Msg) {
		getProvinceQuery := getProvinces.GetProvincesQuery{}
		err := json.Unmarshal(msg.Data, &getProvinceQuery)
		if err != nil {
			log.Println("Error unmarshalling getProvinces ", err)
		}
		newId, _ := uuid.NewUUID()
		provincesResponse := commonResponse.PaginationResponse[responses.ProvinceResponse]{
			Items: []responses.ProvinceResponse{
				{
					Name:          "Ha Noi",
					ModelResponse: commonResponse.ModelResponse{Id: newId.String()},
				},
			},
			CurrentPageIndex: 1,
			TotalPage:        1,
			TotalRecord:      1,
		}
		response, err := json.Marshal(provincesResponse)
		_ = conn.Publish(msg.Reply, response)
	})
	return &NatsSubscriber{conn: conn}
}
