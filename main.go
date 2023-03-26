package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
)

func main() {
	r := gin.Default()
	health := healthcheck.NewHandler()

	health.AddLivenessCheck("ping", healthcheck.HTTPGetCheck("http://localhost:8080", 200*time.Millisecond))
	health.AddReadinessCheck("ping", healthcheck.HTTPGetCheck("http://localhost:8080", 200*time.Millisecond))

	go http.ListenAndServe("0.0.0.0:8086", health)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
