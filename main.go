package main

import (
	"flag"
	"github.com/codegangsta/martini"
)

var Gofi *GoForIt

func main() {
	configFile := flag.String("config", "goforit.json", "Config file")
	flag.Parse()
	Gofi = NewGoForIt()
	Gofi.Init(*configFile)

	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
