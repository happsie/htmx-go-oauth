package service

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	tmpl *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func InitTemplates() *Template {
	return &Template{
		template.Must(template.ParseGlob("tmpl/*.html")),
	}
}
