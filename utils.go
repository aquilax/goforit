package main

import (
	"github.com/aquilax/cyrslug"
	"github.com/aquilax/tripcode"
)

func genTripcode(password string) string {
	return tripcode.Tripcode(password)
}

func slug(text string) string {
	return cyrslug.Slug(text) + ".html"
}
