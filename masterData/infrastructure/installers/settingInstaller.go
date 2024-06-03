package installers

import (
	buildingBlockConfiguration "natsMicros/buildingBlocks/application/configurations"
	"natsMicros/masterData/application/configurations"
	"reflect"
)

func NewMongoDbSettingConfig(configuration *configurations.Configuration) *buildingBlockConfiguration.MongoDbSetting {
	return &configuration.MongoDbSetting
}

func NewMongoDbConnectionName[TModel any]() *buildingBlockConfiguration.MongoDbCollectionSetting[TModel] {
	var model TModel
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		return &buildingBlockConfiguration.MongoDbCollectionSetting[TModel]{CollectionName: t.Elem().Name()}
	}
	return &buildingBlockConfiguration.MongoDbCollectionSetting[TModel]{CollectionName: t.Name()}
}
