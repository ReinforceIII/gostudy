package mainpage

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	r.GET("/test", Hello)

	var err error
	c.Request, err = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(res, c.Request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, "Hello Go!", res.Body.String())
}
