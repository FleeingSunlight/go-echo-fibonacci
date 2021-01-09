package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	e := echo.New()
	e.GET("/:n", func(c echo.Context) error {
		n := c.Param("n")
		i, _ := strconv.Atoi(n)
		return c.JSON(http.StatusOK, fib(i))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
