package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	r.Use(HealthCheck)

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/check", nil)
	r.ServeHTTP(res, c.Request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, "OK", res.Body.String())
}
