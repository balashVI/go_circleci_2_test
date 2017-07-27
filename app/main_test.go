package main_test

import (
	"testing"

	. "github.com/balashVI/go_circleci_2_test/app"
	. "github.com/smartystreets/goconvey/convey"
)

func TextGetHello(t *testing.T) {
	Convey("Test GetHello", t, func() {
		helloMsg := GetHello()
		So(helloMsg, ShouldEqual, "Hello")
	})
}
