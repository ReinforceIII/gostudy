package main

import (
	healthcheck "github.com/ReinforceIII/gostudy/api/healthcheck"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/check", healthcheck.HealthCheck)
	r.Run()
}
