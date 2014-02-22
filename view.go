package main

type PathLink struct {
	Url   string
	Label string
}

type PageData struct {
	Domain *Domain
	Path   *[]PathLink
}
