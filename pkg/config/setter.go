package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error reading config file")
	}

	if err := viper.Unmarshal(&configurations); err != nil {
		fmt.Printf("unable to decode into struct ,%v ", err)
	}
}
