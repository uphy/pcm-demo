package common

import (
	"io"
	"io/fs"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func NewTemplate(filesystem fs.FS) *Template {
	return &Template{
		templates: template.Must(template.ParseFS(filesystem, "*")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
