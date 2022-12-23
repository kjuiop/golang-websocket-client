package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	ProductionMode string `envconfig:"CHAT_CLIENT_PRODUCTION_MODE" default:"local"`

	LogLevel string `envconfig:"CHAT_CLIENT_LOG_LEVEL" default:"debug"`
	LogPath  string `envconfig:"CHAT_CLIENT_LOG_PATH" default:"/home/jake/chat-client/chat-client.log"`

	ApiInfo struct {
		Port string `envconfig:"CHAT_CLIENT_API_PORT" default:"3020"`
	}
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
