package main

import (
	"embed"
	"fmt"
	"pcm-demo/common"

	"github.com/labstack/echo/v4"
)

//go:embed static/*
var static embed.FS

func main() {
	app, err := common.NewApp(static)
	if err != nil {
		panic(err)
	}

	app.Template("/", "index.html")
	app.Template("/index.html", "index.html")
	app.Template("/trigger.html", "trigger.html")
	app.Echo.POST("/.well-known/private-click-measurement/report-attribution/", func(c echo.Context) error {
		fmt.Println("!!!report!!!")
		return nil
	})

	if err := app.Start(80); err != nil {
		panic(err)
	}
}
