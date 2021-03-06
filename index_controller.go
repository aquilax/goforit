package main

import (
	"fmt"
	"github.com/martini-contrib/render"
	"net/http"
)

type IndexPageData struct {
	GroupedBoards *GroupedBoards
	LatestPosts   *[]Post
	PageData
}

func GetHomeIndex(r render.Render, req *http.Request, app *App) {
	indexData := &IndexPageData{}
	indexData.Domain = app.Model.getDomain(req.Host)
	indexData.GroupedBoards = app.Model.getGroupedBoards(indexData.Domain.DomainId)
	indexData.LatestPosts = app.Model.getLatestPosts(indexData.Domain.DomainId, 0, 10)
	indexData.Path = &[]PathLink{}
	fmt.Printf("%#v", indexData)
	r.HTML(200, "index", indexData)
}
