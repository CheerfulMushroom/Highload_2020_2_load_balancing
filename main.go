package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	router := echo.New()
	router.GET("*",
		func(ctx echo.Context) error {
			const maxWaitingTime = 1000
			time.Sleep(time.Duration(rand.Intn(maxWaitingTime)) * time.Millisecond)
			return ctx.String(http.StatusOK, "Hello, I am fast backend server\n")
		},
	)
	log.Fatal(router.Start(":8080"))
}
