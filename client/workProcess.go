package client

import (
	"fmt"
	"golang-websocket-client/form"
	"golang-websocket-client/pkg/logger"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func TestSocketConnectionProcess(request form.SocketConnectionRequest, log *logger.Logger) (error, string) {

	start := time.Now()
	testHost := fmt.Sprintf("%s:%d", request.Host, request.Port)
	prefix := log.TestSocketConnectionProcessPrefix(start)

	u := url.URL{Scheme: request.Protocol, Host: testHost, Path: request.ApiPath}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err, fmt.Sprintf("dial: %s", err)
	}
	defer c.Close()

	time.Sleep(time.Second * time.Duration(request.Period))

	end := time.Now()

	prefix.Data["end"] = end
	prefix.Data["test_host"] = testHost
	prefix.Data["api_path"] = request.ApiPath
	log.Info(prefix, "[TestSocketConnectionProcess] SUCCESS")
	return nil, ""
}
