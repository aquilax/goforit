package main

import (
	"flag"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"time"
)

type GoForIt struct {
}

func NewGoForIt () *GoForIt {
	return &GoForIt{}
}

func (gfi *GoForIt) Run () {
	configFile := flag.String("config", "goforit.json", "Config file")
	flag.Parse()
	app := NewApp()
	app.Init(*configFile)

	m := martini.Classic()
	m.Map(app)
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
		Funcs: []template.FuncMap{
			{
				"lang": func(s string) string {
					// TODO: translate string
					return s
				},
				"mod": func(n int, mod int) int {
					return n % mod
				},
				"slug": func(s string) string {
					return slug(s)
				},
				"time": func(t time.Time) string {
					return t.Format("01.02.2006 15.04")
				},
			},
		},
	}))
	m.Use(martini.Static("public_html"))
	m.Get("/", GetHomeIndex)
	m.Run()

}
