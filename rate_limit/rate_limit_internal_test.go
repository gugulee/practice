package ratelimit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/practice/rate_limit/config"
	"github.com/stretchr/testify/require"
)

func Test_testRateLimit(t *testing.T) {
	path := "/v1/user"
	appId := "app-1"

	ts := httptest.NewServer(setupServer(t, appId, path))
	defer ts.Close()

	resp, err := http.Get(ts.URL + path)
	require.NoError(t, err)

	bs, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	fmt.Println(string(bs))

	require.Equal(t, 200, resp.StatusCode)
}

func setupServer(t *testing.T, appId, path string) *gin.Engine {
	r := gin.Default()
	configParser, err := config.NewConfigParser("yaml")
	require.NoError(t, err)

	ratelimiter := NewRateLimiter(configParser)

	r.GET(path, func(c *gin.Context) {
		if !ratelimiter.Limit(appId, path) {
			c.String(http.StatusBadRequest, "has too much requests")
			return
		}
		pingEndpoint(c)
	})

	return r
}

func pingEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "access granted")
}
