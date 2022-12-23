package client

import (
	"log"
	"os"

	"golang-websocket-client/pkg/config"
	"golang-websocket-client/pkg/logger"
)

type Client struct {
	cfg *config.Configuration

	log *logger.Logger
}

func NewHandler() (*Client, error) {
	c := new(Client)

	// config init
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Println("[main] failed ConfInitialize :", err)
		os.Exit(1)
	}

	clientLogger, err := logger.LogInitialize(cfg.LogLevel, cfg.LogPath)
	if err != nil {
		log.Println("[main] failed log Initialize : ", err)
	}

	c.cfg = cfg
	c.log = clientLogger

	return c, nil
}

func (s *Client) Close() {
	log.Println("[main] Client close")
}
