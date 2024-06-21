package services

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"natsMicros/buildingBlocks/infrastructure/helpers/reflection/messageHelper"
)

type NatResponseService[TRequest any, TResponse any] struct {
	natConn *nats.Conn
}

func NewNatResponseService[TRequest any, TResponse any](natConn *nats.Conn) *NatResponseService[TRequest, TResponse] {
	return &NatResponseService[TRequest, TResponse]{natConn}
}

func (natService *NatResponseService[TRequest, TResponse]) Response(cb func(TRequest) (TResponse, error)) error {
	natConn := natService.natConn
	msgChannel := messageHelper.GetMessageExchange[TRequest]()
	_, err := natConn.Subscribe(msgChannel, func(msg *nats.Msg) {
		var msgData TRequest
		err := json.Unmarshal(msg.Data, &msgData)
		var response TResponse
		if err != nil {
			resByte, _ := json.Marshal(err)
			_ = natConn.Publish(msg.Reply, resByte)
			return
		}
		response, err = cb(msgData)
		if err != nil {
			resByte, _ := json.Marshal(err)
			_ = natConn.Publish(msg.Reply, resByte)
			return
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			resByte, _ := json.Marshal(err)
			_ = natConn.Publish(msg.Reply, resByte)
			return
		}
		_ = natConn.Publish(msg.Reply, responseBytes)
	})
	return err
}
