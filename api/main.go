package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world  ")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":8080")))
}
