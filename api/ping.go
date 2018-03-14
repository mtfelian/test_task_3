package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func extractBody(c *gin.Context, to interface{}) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, to)
}

// Ping is an API ping-pong method
func Ping(c *gin.Context) {
	time.Sleep(200 * time.Millisecond)
	c.JSON(http.StatusOK, "pong")
}
