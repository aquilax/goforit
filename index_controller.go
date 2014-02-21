package main

import (
	"net/http"
	"github.com/martini-contrib/render"
)

func GetHomeIndex(r render.Render, req *http.Request, gfi *GoForIt) {
	gfi.Model.getDomain(req.Host)
	gfi.Model.getBoards()
	gfi.Model.getLatestPosts()
	r.HTML(200, "index", "world")
}
