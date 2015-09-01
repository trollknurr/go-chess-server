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

	if from_x != to_x && from_y == to_y {
		return false, MoveError{"Неправильный ход: пешка не ходит вбок"}
	}

	switch {
	case target_cell.color == FREE:
		if from_x != to_x {
			return false, MoveError{"Неправильный ход: пешка по диагонали"}
		}
		if (from_y == 1 && to_y > 3) || (from_y == 6 && to_y < 3) {
			return false, MoveError{"Неправильный ход: пешка максимум на две клетки"}
		}
		field[from_y][from_x] = Cell{EMPTY, FREE}
		field[to_y][to_x] = from_cell
		return true, nil
	case target_cell.color == from_cell.color:
		return false, MoveError{"Неправильный ход: ход на клетку, занятую своей фигурой"}
	case target_cell.color != from_cell.color:
		if from_cell.color == WHITE {
			if ((from_x+1 == to_x) && (from_y+1 == to_y)) || (from_x-1 == to_x) && (from_y+1 == to_y) {
				field[from_y][from_x] = Cell{EMPTY, FREE}
				field[to_y][to_x] = from_cell
				return true, nil
			}
		}
		if from_cell.color == BLACK {
			if ((from_x-1 == to_x) && (from_y-1 == to_y)) || (from_x+1 == to_x) && (from_y-1 == to_y) {
				field[from_y][from_x] = Cell{EMPTY, FREE}
				field[to_y][to_x] = from_cell
				return true, nil
			}
		}
		return false, MoveError{"Неправильный ход: эту фигуру нельзя съесть"}
	}
	return false, MoveError{"Неизвестная ошибка"}
}

func main() {
	f := InitGameField()
	fmt.Println(f[6][1])
}
