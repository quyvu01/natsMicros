package services

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"reflect"
)

type NatPublisherService[TMessage any] struct {
	natConn *nats.Conn
}

func NewNatPublisherService[TMessage any](natsConn *nats.Conn) *NatPublisherService[TMessage] {
	return &NatPublisherService[TMessage]{natConn: natsConn}
}
func (natService *NatPublisherService[TRequest]) Publish(msg TRequest) error {
	natConn := natService.natConn
	request, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	msgChannel := reflect.TypeOf(msg).Name()
	err = natConn.Publish(msgChannel, request)
	return err
}
