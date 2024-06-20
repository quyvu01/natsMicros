package main

import (
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	buildingBlockConfiguration "natsMicros/buildingBlocks/application/configurations"
	"natsMicros/masterData/application/configurations"
	"natsMicros/masterData/domain"
	"natsMicros/masterData/infrastructure/installers"
)

func RunApplication(_ *configurations.Configuration, _ *nats.Conn, _ *installers.NatsSubscriber, _ *mongo.Client, _ *buildingBlockConfiguration.MongoDbCollectionSetting[domain.Province]) {

}
