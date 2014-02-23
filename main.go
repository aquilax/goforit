package main

import (
	"flag"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"time"
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
					// TODO: translate string
					return s
				},
				"mod": func(n int, mod int) int {
					return n % mod
				},
				"slug": func(s string) string {
					// TODO: generate slug
					return s
				},
				"time": func(t time.Time) string {
					// TODO: format time
					return "123"
				},
			},
		},
	}))
	m.Use(martini.Static("public_html"))
	m.Get("/", GetHomeIndex)
	m.Run()
}
