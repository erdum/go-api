package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Name						string `default:"GO-API"`
	Secret						string
	Debug    					bool   `default:"true"`
	Port     					string `default:"8000"`
	Url      					string `default:"http://localhost"`
	Firebase struct {
		ProjectId   			string `split_words:"true"`
		Credentials 			string
	}
	PasswordResetExpirySecs		uint `default:"300" split_words:"true"`
	Otp struct {
		ExpirySecs				uint `default:"20" split_words:"true"`
		RetrySecs				uint `default:"300" split_words:"true"`
		Retries					uint `default:"3" split_words:"true"`
	}
	Mail struct {
		Host					string
		Port					string `default:"465"`
		User					string
		Pass					string
		From					string
	}
}

var appConfig = &Config{}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	if err := envconfig.Process("", appConfig); err != nil {
		return nil, err
	}

	return appConfig, nil
}

func GetConfig() *Config {
	return appConfig
}
