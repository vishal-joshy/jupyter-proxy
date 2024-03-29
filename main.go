package main

import (
	"jupy/jupyter"

	"github.com/labstack/echo/v4"
)

type Header struct {
	Authorization string
}

func getJupyterClient() {

}

func main() {
	e := echo.New()
	e.GET("/users", jupyter.GetUsers)
	e.GET("/users/:name", jupyter.GetUser)
	e.POST("/users", jupyter.CreateUser)
	e.Logger.Fatal(e.Start(":1325"))
}
