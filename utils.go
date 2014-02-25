package main

import (
	"github.com/aquilax/tripcode"
	"github.com/aquilax/cyrslug"
)

func genTripcode(password string) string {
	return tripcode.Tripcode(password)
}

func slug(text string) string {
	return cyrslug.Slug(text) + ".html"
}
