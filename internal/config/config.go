package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database DBConfig       `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbit"`
}

type ServerConfig struct {
	Port    string `mapstructure:"port"`
	Address string `mapstructure:"address"`
}

type DBConfig struct {
	Port     int    `mapstructure:"port"`
	Addr     string `mapstructure:"address"`
	DBName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type RabbitMQConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Queue    string `mapstructure:"queue"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		return &Config{}, err
	}

	return &conf, nil
}
