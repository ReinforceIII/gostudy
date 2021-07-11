package mainpage

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Go!")
}

func execRoutine() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 3)
	}()

	wg.Wait()
}

func RoutineHandler(c *gin.Context) {
	c.String(http.StatusOK, "RoutineEnd")
	execRoutine()
}
