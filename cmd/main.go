package main

import (
	"github.com/ReinforceIII/gostudy/api/healthcheck"
	"github.com/ReinforceIII/gostudy/api/mainpage"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/check", healthcheck.HealthCheck)
	r.GET("/hello", mainpage.Hello)
	r.Run()
}
