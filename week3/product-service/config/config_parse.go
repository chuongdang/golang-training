package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	env := os.Getenv("GO_ENV")
	if env == "prod" {
		viper.SetConfigName(fmt.Sprintf("config_%s", env))
		if err := viper.MergeInConfig(); err != nil {
			log.Fatalf("Fatal error config file: %s \n", err)
		}
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Could not unmarshal")
	}
}
