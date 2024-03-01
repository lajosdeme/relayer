package core

import (
	"fmt"

	"github.com/lajosdeme/transaction-relayer/types"
	"github.com/spf13/viper"
)

var Config types.Config

func Load() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to read env file: ", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		fmt.Println("Failed to decode env file: ", err)
	}
}
