package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck  Server Status Check Handler
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
