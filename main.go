package main

import (
	"github.com/codegangsta/martini"
)

var Gofi *GoForIt

func main() {
	Gofi = NewGoForIt()
	Gofi.Init("")

	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
