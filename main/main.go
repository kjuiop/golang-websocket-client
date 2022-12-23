package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-websocket-client/client"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("hello golang-websocket-client")

	sigs := make(chan os.Signal, 1)
	defer close(sigs)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	c, err := client.NewHandler()
	if err != nil {
		log.Println("[main] failed NewClient :", err)
		os.Exit(1)
	}
	defer c.Close()

	gMux := gin.Default()
	gMux.GET("/api/health-check", c.HealthCheck)
	gMux.POST("/api/test/socket-connection", c.TestSocketConnection)

	apiPort := c.GetApiPort()
	gMux.Run(":" + apiPort)

	/** worker 형태로 사용 시 이용
	wg := sync.WaitGroup{}

	wg.Add(1)
	go c.CloseWithContext(sigs, &wg)

	wg.Wait()
	*/

}
