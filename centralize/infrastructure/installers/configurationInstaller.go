package installers

import (
	"github.com/spf13/viper"
	"natsMicros/centralize/application/configurations"
)

func NewConfiguration() *configurations.Configuration {
	configuration := configurations.Configuration{}
	viper.SetConfigFile("./centralize/api/appsettings.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Can't find the file appsettings.json")
	}
	err = viper.Unmarshal(&configuration)
	if err != nil {
		panic("Environment can't be loaded!")
	}
	return &configuration
}
