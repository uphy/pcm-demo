package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
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
	app.Echo.GET("/tracking-pixel.gif", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/.well-known/private-click-measurement/trigger-attribution/10")
	})
	app.Echo.GET("/.well-known/private-click-measurement/trigger-attribution/:triggerData/:priority", func(c echo.Context) error {
		log.Println("Trigger attribution with priority")
		return c.Blob(200, "image/gif", []byte{})
	})
	app.Echo.GET("/.well-known/private-click-measurement/trigger-attribution/:triggerData", func(c echo.Context) error {
		log.Println("Trigger attribution")
		return c.Blob(200, "image/gif", []byte{})
	})
	app.Echo.POST("/.well-known/private-click-measurement/report-attribution/", func(c echo.Context) error {
		fmt.Println("!!!report!!!")
		return nil
	})

	if err := app.Start(80); err != nil {
		panic(err)
	}
}
