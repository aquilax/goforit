package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoad(t *testing.T) {
	Convey("Given config", t, func() {
		config := NewConfig()
		Convey("When file name is wrong", func() {
			err := config.Load("")
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, "open : no such file or directory")
		})
	})
}
