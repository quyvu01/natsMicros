package services

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"natsMicros/buildingBlocks/infrastructure/helpers/reflection/messageHelper"
)

type NatConsumerService[TMessage any] struct {
	natConn *nats.Conn
}

func NewNatConsumerService[TMessage any]() (*NatConsumerService[TMessage], error) {
	return &NatConsumerService[TMessage]{natConn: nil}, nil
}

func (natService *NatConsumerService[TMessage]) Consume(cb func(TMessage, error)) error {
	natConn := natService.natConn
	msgChannel := messageHelper.GetMessageExchange[TMessage]()
	_, err := natConn.Subscribe(msgChannel, func(msg *nats.Msg) {
		var msgData TMessage
		err := json.Unmarshal(msg.Data, &msgData)
		cb(msgData, err)
	})
	return err
}
