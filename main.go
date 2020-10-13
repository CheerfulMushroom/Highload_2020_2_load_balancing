package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var hitsCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "hits_counter",
	Help: "The total number of hits",
})

func ServeContent(ctx echo.Context) error {
	const maxWaitingTime = 1000
	time.Sleep(time.Duration(rand.Intn(maxWaitingTime)) * time.Millisecond)
	return ctx.String(http.StatusOK, "Hello, I am fast backend server\n")
}


func IncreaseHits(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hitsCounter.Inc()
		return next(c)
	}
}

func main() {

	go func() {
		prometheusRouter := echo.New()
		prometheusRouter.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		log.Fatal(prometheusRouter.Start(":5050"))
	}()

	router := echo.New()
	router.Use(IncreaseHits)
	router.GET("/", ServeContent,
	)
	log.Fatal(router.Start(":8080"))
}
