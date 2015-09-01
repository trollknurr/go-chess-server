package main

import (
	"fmt"
)

type MoveError struct {
	message string
}

func (e MoveError) Error() string {
	return e.message
}

func (field *Field) Move(from_x, from_y, to_x, to_y int) (bool, error) {
	unit := field[from_y][from_x].unit
	switch {
	case unit == MAN:
		return field.move_man(from_x, from_y, to_x, to_y)
	}
	return false, nil
}

func (field *Field) move_man(from_x, from_y, to_x, to_y int) (bool, error) {
	from_cell := field[from_y][from_x]
	target_cell := field[to_y][to_x]

	switch {
	case from_cell.color == WHITE && target_cell.color == FREE:
		if from_x != to_x {
			return false, MoveError{"Неправильный ход: пешка по диагонали"}
		}
		if from_y == 1 && to_y > 3 {
			return false, MoveError{"Неправильный ход: пешка максимум на две клетки"}
		}
	}
	return true, nil
}

func main() {
	f := InitGameField()
	res, err := f.Move(1, 1, 2, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
