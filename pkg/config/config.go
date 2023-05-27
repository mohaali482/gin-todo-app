package config

import (
	"github.com/spf13/viper"
)

type Config struct {
    Port  string `mapstructure:"PORT"`
    DB_URL string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
    viper.AddConfigPath("./")
    viper.SetConfigName(".env")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()

    if err != nil {
        return Config{}, err
    }

    err = viper.Unmarshal(&c)

    return c, nil
}