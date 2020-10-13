package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	router := echo.New()
	router.GET("*",
		func(ctx echo.Context) error {
			return ctx.String(http.StatusOK, "Hello, I am fast backend server\n")
		},
	)
	log.Fatal(router.Start(":8080"))
}
