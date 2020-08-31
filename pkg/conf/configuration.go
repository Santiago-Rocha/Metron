package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database DatabaseConfiguration
}

var configuration Configuration

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("pkg/conf/")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file %s", err)

	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logrus.Fatalf("Unable to decode into struct %v", err)
	}
}

func GetConfiguration() Configuration {
	return configuration
}
