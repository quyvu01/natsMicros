package services

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"natsMicros/buildingBlocks/infrastructure/helpers/reflection/messageHelper"
	"time"
)

const defaultTimeout = 10 * time.Second

type NatRequestService[TRequest any, TResponse any] struct {
	natConn *nats.Conn
}

func NewNatRequestService[TRequest any, TResponse any](natConn *nats.Conn) *NatRequestService[TRequest, TResponse] {
	return &NatRequestService[TRequest, TResponse]{natConn}
}

func (natService *NatRequestService[TRequest, TResponse]) Request(msg TRequest) (TResponse, error) {
	natConn := natService.natConn
	request, err := json.Marshal(msg)
	response := new(TResponse)
	if err != nil {
		return *response, err
	}
	msgChannel := messageHelper.GetMessageExchange[TRequest]()
	res, err := natConn.Request(msgChannel, request, defaultTimeout)
	if err != nil {
		return *response, err
	}
	err = json.Unmarshal(res.Data, response)
	if err != nil {
		return *response, err
	}
	return *response, nil
}
