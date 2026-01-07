package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("crud_api/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("../../..")
	viper.AddConfigPath("../../../..")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// fmt.Println("debug", viper.GetString("DB_HOST"))
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	// fmt.Println("ini cfg", cfg)
	return cfg
}
