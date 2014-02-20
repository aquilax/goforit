package main

import (
	"github.com/aquilax/tripcode"
)

func genTripcode(password string) string {
	return tripcode.Tripcode(password)
}
