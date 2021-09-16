package ratelimit

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func Test_NewRateLimiter(t *testing.T) {
	_, err := NewRateLimiter("yaml", "config/config.yaml", "RollingWindow")
	require.NoError(t, err)
}

func Test_RateLimit(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	ratelimiter, err := NewRateLimiter("yaml", "config/config.yaml", "RollingWindow")
	require.NoError(t, err)

	ts := httptest.NewServer(setupServer(ratelimiter))
	defer ts.Close()
	time.Sleep(1 * time.Hour)

	/*
		TODO: run test in shell:
		while true;
		do
		curl $(ss -tlnp | grep rate_limit | awk '{print $4}')/v1/user; echo
		done
	*/

}

func setupServer(ratelimiter *RateLimiter) *gin.Engine {
	r := gin.Default()

	app:="app-1"
	path:="/v1/user"
	r.GET(path, ratelimite(pingEndpoint, ratelimiter, app, path))

	path="/v1/order"
	r.GET(path, ratelimite(pingEndpoint, ratelimiter, app, path))

	return r
}

func ratelimite(handler gin.HandlerFunc, ratelimiter *RateLimiter, appId, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ratelimiter.Limit(appId, path) {
			c.String(http.StatusBadRequest, "too much request")
			return
		}
		handler(c)
	}
}

func pingEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "access granted")
}
