package client

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-websocket-client/form"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"golang-websocket-client/pkg/config"
	"golang-websocket-client/pkg/logger"
)

type Client struct {
	cfg *config.Configuration

	log *logger.Logger

	gCtx    context.Context
	gCancel context.CancelFunc
}

func NewHandler() (*Client, error) {
	c := new(Client)

	ctx, cancel := context.WithCancel(context.Background())
	c.gCtx = ctx
	c.gCancel = cancel

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

func (c *Client) Close() {
	log.Println("[main] Client close")
}

func (c *Client) HealthCheck(gCtx *gin.Context) {

	gCtx.JSON(http.StatusOK, map[string]string{"result": "success"})
	return
}

func (c *Client) TestSocketConnection(gCtx *gin.Context) {

	prefix := c.log.InitPrefixData()

	var request form.SocketConnectionRequest
	var response form.SocketConnectionResponse

	err, errorMsg := c.readBody(gCtx, &request)
	if err != nil {
		c.log.Error(prefix, errorMsg)
		response.Message = errorMsg
		gCtx.Status(http.StatusBadRequest)
		return
	}

	prefix.Data["request_param"] = request
	c.log.Info(prefix, "[TEST-SocketConnection] API TEST")

	err, errorMsg = TestSocketConnectionProcess(request, c.log)
	if err != nil {
		c.log.Error(prefix, errorMsg)
		response.Message = errorMsg
		gCtx.Status(http.StatusInternalServerError)
		return
	}

	response.Result = "success"
	gCtx.JSON(http.StatusOK, response)
	return
}

func (c *Client) GetApiPort() string {
	return c.cfg.ApiInfo.Port
}

func (c *Client) CloseWithContext(sigs chan os.Signal, wg *sync.WaitGroup) {

	prefix := c.log.InitPrefixData()

	for {
		select {
		case <-sigs:
			c.log.Debug(prefix, "Receive exit signal")
			c.gCancel()
		case <-c.gCtx.Done():
			c.log.Debug(prefix, "CloseWithContext Close Goroutine")
			wg.Done()
			return
		default:
			time.Sleep(time.Second * 1)
		}
	}
}
