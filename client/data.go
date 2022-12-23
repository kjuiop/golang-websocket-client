package client

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// readBody parses the json body of the http message.
func (c *Client) readBody(gCtx *gin.Context, data interface{}) (error, string) {

	b, err := ioutil.ReadAll(gCtx.Request.Body)
	if err != nil {
		message := fmt.Sprintf("[readBody] Body Parse Error : %s", err.Error())
		return err, message
	}
	if len(b) == 0 {
		message := fmt.Sprintf("[readBody] Empty Parameters : %s", string(b))
		return fmt.Errorf("empty Parameter"), message
	}

	err = json.Unmarshal(b, data)
	if err != nil {
		message := fmt.Sprintf("[readBody] Body Parse Error : %s", err.Error())
		return err, message
	}

	return nil, ""
}
