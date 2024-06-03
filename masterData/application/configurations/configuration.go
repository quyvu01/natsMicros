package configurations

import "natsMicros/buildingBlocks/application/configurations"

type Configuration struct {
	NatsSetting    NatsSSetting
	MongoDbSetting configurations.MongoDbSetting
}
