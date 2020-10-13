package Highload_2020_2_load_balancing

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	router := echo.New()
	router.GET("/api",
		func(ctx echo.Context) error {
			return ctx.String(http.StatusOK, "hello")
		},
	)
	log.Fatal(router.Start(":8080"))
}
