package main

import (
	"flag"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func main() {
	configFile := flag.String("config", "goforit.json", "Config file")
	flag.Parse()
	Gofi := NewGoForIt()
	Gofi.Init(*configFile)

	m := martini.Classic()
	m.Map(Gofi)
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(martini.Static("public_html"))
	m.Get("/", GetHomeIndex)
	m.Run()
}
