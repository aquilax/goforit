package main

import (
	"flag"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"html/template"
)

func main() {
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
					return s
				},
			},
		},
	}))
	m.Use(martini.Static("public_html"))
	m.Get("/", GetHomeIndex)
	m.Run()
}
