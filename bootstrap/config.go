package bootstrap

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Loads configs with the provided names.
func LoadConfigs(configs ...string) {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")

	for _, config := range configs {
		viper.SetConfigName(config)

		if err := viper.MergeInConfig(); err != nil {
			logrus.WithError(err).Fatalln("unable to load configuration:", config)
		}
	}
}
