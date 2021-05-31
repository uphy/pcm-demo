package common

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	App struct {
		Echo              *echo.Echo
		TemplateVariables *TemplateVariables
	}
	TemplateVariables struct {
		NgrokEndpoints *NgrokEndpoints
	}
)

func NewApp(static fs.FS) (*App, error) {
	public, err := fs.Sub(static, "static/public")
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.Renderer = NewTemplate(public)
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	ngrokEndpoints, err := WaitForNgrok()
	if err != nil {
		return nil, err
	}
	templateVariables := TemplateVariables{
		NgrokEndpoints: ngrokEndpoints,
	}

	log.Println("Media: ", ngrokEndpoints.Media)
	log.Println("Advertiser: ", ngrokEndpoints.Advertiser)

	app := &App{e, &templateVariables}
	app.Echo.GET("/favicon.ico", func(c echo.Context) error {
		return c.Blob(200, "image/vnd.microsoft.icon", []byte{})
	})
	return app, nil
}

func (a *App) Template(path string, templateName string) {
	a.Echo.GET(path, func(c echo.Context) error {
		return c.Render(200, templateName, a.TemplateVariables)
	})
}

func (a *App) Start(port int) error {
	return a.Echo.Start(fmt.Sprintf(":%d", port))
}
