package client

import (
	"golang-websocket-client/pkg/config"
	"log"
	"os"
)

type Client struct {
	cfg *config.Configuration
}

func NewHandler() (*Client, error) {
	c := new(Client)

	// config init
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Println("[main] failed ConfInitialize :", err)
		os.Exit(1)
	}

	c.cfg = cfg

	return c, nil
}

func (s *Client) Close() {
	log.Println("[main] Client close")
}
