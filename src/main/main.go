package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the website")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.String(http.StatusOK, fmt.Sprintf("your catName is: %s\nand his type is: %s", catName, catType))
}

func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()

	e.GET("/",yallo)

	e.GET("/cats", getCats)

	e.Start(":1323")

}
