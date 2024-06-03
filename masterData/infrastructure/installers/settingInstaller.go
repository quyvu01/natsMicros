package installers

import (
	buildingBlockConfiguration "natsMicros/buildingBlocks/application/configurations"
	"natsMicros/masterData/application/configurations"
)

func NewMongoDbSettingConfig(configuration *configurations.Configuration) *buildingBlockConfiguration.MongoDbSetting {
	return &configuration.MongoDbSetting
}
