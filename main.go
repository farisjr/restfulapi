package main

import (
	"fmt"
	"test/project/configs"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	configs.InitDb()
	configs.InitPort()
	//middlewares.LogMiddlewares(e)
	///routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", configs.HTTP_PORT)))
}
