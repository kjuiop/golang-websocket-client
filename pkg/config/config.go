package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	ProductionMode string `envconfig:"CHAT_PRODUCTION_MODE" default:"local"`
}

// 환경 변수 설정
func init() {
	println("configuration init")
}

//LoadEnv loads configuration from environment variables
func LoadEnv() (*Configuration, error) {
	var config Configuration
	err := envconfig.Process("golang-websocket-client", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
