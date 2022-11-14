package main

import (
	"github.com/labstack/echo"
)

func main() {
	// create echo instance
	e := echo.New()
	e.Logger.Fatal(e.Start(":8000"))
}