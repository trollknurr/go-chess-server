package main

type MoveError struct {
	message string
}

func (e MoveError) Error() string {
	return e.message
}

func (field *Field) Move(from_x, from_y, to_x, to_y int) (bool, error) {
	unit := field[from_y][from_x].CellUnit
	switch {
	case unit == MAN:
		valid_move, err := field.validate_man_move(from_x, from_y, to_x, to_y)
		if valid_move && err == nil {
			field.move_figure(from_x, from_y, to_x, to_y)
		}
		return valid_move, err
	}
	return false, nil
}

func (field *Field) getCell(x, y int) Cell {
	return field[y][x]
}

func (field *Field) move_figure(from_x, from_y, to_x, to_y int) {
	from_cell := field[from_y][from_x]
	field[from_y][from_x] = Cell{EMPTY, FREE}
	field[to_y][to_x] = from_cell
}

func (field *Field) validate_man_move(from_x, from_y, to_x, to_y int) (bool, error) {
	from_cell := field[from_y][from_x]
	target_cell := field[to_y][to_x]

	if from_x != to_x && from_y == to_y {
		return false, MoveError{"Неправильный ход: пешка не ходит вбок"}
	}
	switch {
	case target_cell.CellColor == FREE:
		if from_x != to_x {
			return false, MoveError{"Неправильный ход: пешка по диагонали"}
		}
		if (from_y == 1 && to_y > 3) || (from_y == 6 && to_y < 3) {
			return false, MoveError{"Неправильный ход: пешка максимум на две клетки"}
		}
		return true, nil
	case target_cell.CellColor == from_cell.CellColor:
		return false, MoveError{"Неправильный ход: ход на клетку, занятую своей фигурой"}
	case target_cell.CellColor != from_cell.CellColor:
		if from_cell.CellColor == WHITE {
			if ((from_x+1 == to_x) && (from_y+1 == to_y)) || (from_x-1 == to_x) && (from_y+1 == to_y) {
				return true, nil
			}
		}
		if from_cell.CellColor == BLACK {
			if ((from_x-1 == to_x) && (from_y-1 == to_y)) || (from_x+1 == to_x) && (from_y-1 == to_y) {
				return true, nil
			}
		}
		return false, MoveError{"Неправильный ход: эту фигуру нельзя съесть"}
	}
	return false, MoveError{"Неизвестная ошибка"}
}
