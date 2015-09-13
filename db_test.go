package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDBInteraction(t *testing.T) {
	SetDefaultFailureMode(FailureContinues)
	defer SetDefaultFailureMode(FailureHalts)
	Convey("Test database interaction", t, func() {
		session := GetSession()

		Convey("Test new game", func() {
			collection := session.DB("test_chess").C("test_games")
			game := MakeDBNewGame(collection)

			Convey("New game init", func() {
				So(game, ShouldHaveSameTypeAs, Game{})
			})

			Convey("Find game in DB", func() {
				query := collection.FindId(game.Id)
				value, _ := query.Count()
				So(value, ShouldEqual, 1)
			})

			Convey("Fetch game field from DB", func() {
				f_game := GetGameById(game.Id, collection)

				Convey("Field created", func() {
					f := f_game.Field
					So(len(f), ShouldEqual, 8)
					So(len(f[0]), ShouldEqual, 8)
				})

				Convey("Cell fetched", func() {
					f := f_game.Field
					cell := f.getCell(0, 0)
					So(cell, ShouldHaveSameTypeAs, Cell{})
					So(cell.CellColor, ShouldEqual, WHITE)
					So(cell.CellUnit, ShouldEqual, CASTLE)
				})

				Convey("Make move and save", func() {
					f_game.Field.Move(0, 1, 0, 3)
					SaveGameById(f_game, collection)
					another_game := GetGameById(f_game.Id, collection)

					Convey("To cell", func() {
						cell := another_game.Field.getCell(0, 3)
						So(cell.CellUnit, ShouldEqual, MAN)
						So(cell.CellColor, ShouldEqual, WHITE)
					})

					Convey("From cell", func() {
						cell := another_game.Field.getCell(0, 1)
						So(cell.CellUnit, ShouldEqual, EMPTY)
						So(cell.CellColor, ShouldEqual, FREE)
					})

				})
			})

			// session.DB("test_chess").DropDatabase()
		})
	})
}
