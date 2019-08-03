package main

import (
	"github.com/ashishthakur913/project/db"
	"github.com/ashishthakur913/project/handler"
	"github.com/ashishthakur913/project/router"
	"github.com/ashishthakur913/project/store"
	"github.com/ashishthakur913/project/apphtml"
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	r := router.New()
	v1 := r.Group("/api")

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	r.Renderer = t
	r.GET("/home", apphtml.HomePage)

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewArticleStore(d)
	h := handler.NewHandler(us, as)
	h.Register(v1)
	r.Logger.Fatal(r.Start(":80"))
}
