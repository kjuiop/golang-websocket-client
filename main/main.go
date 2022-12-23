package main

import (
	"fmt"
	"golang-websocket-client/client"
	"log"
	"os"
)

func main() {
	fmt.Println("hello golang-websocket-client")

	c, err := client.NewHandler()
	if err != nil {
		log.Println("[main] failed NewScheduler :", err)
		os.Exit(1)
	}
	defer c.Close()
}
