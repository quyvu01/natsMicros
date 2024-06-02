package installers

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	responses2 "natsMicros/buildingBlocks/application/responses"
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
		provinceResponse := responses.ProvinceResponse{
			Name:          "Ha Noi",
			ModelResponse: responses2.ModelResponse{Id: "1"},
		}
		response, err := json.Marshal(provinceResponse)
		_ = conn.Publish(msg.Reply, response)
	})
	return &NatsSubscriber{conn: conn}
}
