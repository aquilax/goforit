package main

import (
	"github.com/martini-contrib/render"
	"net/http"
)

func GetHomeIndex(r render.Render, req *http.Request, app *App) {
	app.Model.getDomain(req.Host)
	app.Model.getBoards()
	app.Model.getLatestPosts()
	r.HTML(200, "index", "world")
}
