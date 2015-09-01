package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUnitMovesRight(t *testing.T) {
	SetDefaultFailureMode(FailureContinues)
	defer SetDefaultFailureMode(FailureHalts)
	Convey("Test moves from start position", t, func() {
		field := InitGameField()

		Convey("Test white MAN not moves to the right", func() {
			res, _ := field.Move(1, 1, 2, 1)
			So(res, ShouldBeFalse)
		})
		Convey("Test white MAN not moves to the left", func() {
			res, _ := field.Move(1, 1, 0, 1)
			So(res, ShouldBeFalse)
		})
		Convey("Test white MAN moves forward", func() {
			res, _ := field.Move(1, 1, 1, 2)
			So(res, ShouldBeTrue)
			So(field[2][1].color, ShouldEqual, WHITE)
			So(field[2][1].unit, ShouldEqual, MAN)
		})
		Convey("Test black MAN not moves to the right", func() {
			res, _ := field.Move(1, 6, 2, 6)
			So(res, ShouldBeFalse)
		})
		Convey("Test black MAN not moves to the left", func() {
			res, _ := field.Move(1, 6, 0, 6)
			So(res, ShouldBeFalse)
		})
		Convey("Test black MAN moves forward", func() {
			So(field[6][1].color, ShouldEqual, BLACK)
			So(field[6][1].unit, ShouldEqual, MAN)
			res, _ := field.Move(1, 6, 1, 4)
			So(res, ShouldBeTrue)
			So(field[4][1].color, ShouldEqual, BLACK)
			So(field[4][1].unit, ShouldEqual, MAN)
		})
	})

}
