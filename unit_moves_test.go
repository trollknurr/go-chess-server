package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUnitMovesRight(t *testing.T) {
	SetDefaultFailureMode(FailureContinues)
	defer SetDefaultFailureMode(FailureHalts)
	Convey("Test moves from start position", t, func() {

		Convey("Test man moves", func() {

			field := InitGameField()

			Convey("Test white MAN not moves to the right", func() {
				res, _ := field.Move(1, 1, 2, 1)
				So(res, ShouldBeFalse)
			})
			Convey("Test white MAN not moves on empty diag", func() {
				res, _ := field.Move(1, 1, 2, 2)
				So(res, ShouldBeFalse)
			})
			Convey("Test white MAN not moves to the left", func() {
				res, _ := field.Move(1, 1, 0, 1)
				So(res, ShouldBeFalse)
			})
			Convey("Test white MAN moves forward", func() {
				res, _ := field.Move(1, 1, 1, 2)
				So(res, ShouldBeTrue)
				So(field.getCell(1, 2).CellColor, ShouldEqual, WHITE)
				So(field.getCell(1, 2).CellUnit, ShouldEqual, MAN)
			})
			Convey("Test black MAN not moves to the right", func() {
				res, _ := field.Move(1, 6, 2, 6)
				So(res, ShouldBeFalse)
			})
			Convey("Test black MAN not moves on empty diag", func() {
				res, _ := field.Move(1, 6, 2, 5)
				So(res, ShouldBeFalse)
			})
			Convey("Test black MAN not moves to the left", func() {
				res, _ := field.Move(1, 6, 0, 6)
				So(res, ShouldBeFalse)
			})
			Convey("Test black MAN moves forward", func() {
				res, _ := field.Move(1, 6, 1, 4)
				So(res, ShouldBeTrue)
				So(field.getCell(1, 4).CellColor, ShouldEqual, BLACK)
				So(field.getCell(1, 4).CellUnit, ShouldEqual, MAN)
			})
		})
	})
	Convey("Test pieces eat", t, func() {

	})

}
