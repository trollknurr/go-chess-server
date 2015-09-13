package main

import (
	"fmt"
	"strings"
)

// Фигуры
type Unit int

const (
	EMPTY Unit = 0 + iota
	MAN
	CASTLE
	KNIGHT
	BISHOP
	QUEEN
	KING
)

var units = [...]string{
	"EMPTY",
	"MAN",
	"CASTLE",
	"KNIGHT",
	"BISHOP",
	"QUEEN",
	"KING",
}

func (m Unit) String() string { return units[m] }

// Цвет фигуры
type Color int

const (
	FREE Color = 0 + iota
	WHITE
	BLACK
)

var colors = [...]string{
	"FREE",
	"WHITE",
	"BLACK",
}

func (m Color) String() string { return colors[m] }

// Клетка
type Cell struct {
	CellUnit  Unit
	CellColor Color
}

var cellfigures = [...]string{
	" ",
	"♙",
	"♖",
	"♘",
	"♗",
	"♕",
	"♔",
	"♟",
	"♜",
	"♞",
	"♝",
	"♛",
	"♚",
}

func (c Cell) String() string {
	i := c.CellUnit
	if c.CellColor == BLACK {
		i += 6
	}
	return fmt.Sprintf("%s", cellfigures[i])
}

//Поле
type Field [8][8]Cell

func (f Field) String() string {
	rows := [8]string{}
	row := [8]string{}
	for i := range f {
		for j := range f[i] {
			row[j] = f[i][j].String()
		}
		rows[i] = strings.Join(row[:], " ")
	}
	return fmt.Sprintf("%s\n", strings.Join(rows[:], "\n"))
}
func InitGameField() Field {
	return Field{
		{Cell{CASTLE, WHITE}, Cell{KNIGHT, WHITE}, Cell{BISHOP, WHITE}, Cell{QUEEN, WHITE}, Cell{KING, WHITE}, Cell{BISHOP, WHITE}, Cell{KNIGHT, WHITE}, Cell{CASTLE, WHITE}},
		{Cell{MAN, WHITE}, Cell{MAN, WHITE}, Cell{MAN, WHITE}, Cell{MAN, WHITE}, Cell{MAN, WHITE}, Cell{MAN, WHITE}, Cell{MAN, WHITE}, Cell{MAN, WHITE}},
		{Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}},
		{Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}},
		{Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}},
		{Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}, Cell{EMPTY, FREE}},
		{Cell{MAN, BLACK}, Cell{MAN, BLACK}, Cell{MAN, BLACK}, Cell{MAN, BLACK}, Cell{MAN, BLACK}, Cell{MAN, BLACK}, Cell{MAN, BLACK}, Cell{MAN, BLACK}},
		{Cell{CASTLE, BLACK}, Cell{KNIGHT, BLACK}, Cell{BISHOP, BLACK}, Cell{KING, BLACK}, Cell{QUEEN, BLACK}, Cell{BISHOP, BLACK}, Cell{KNIGHT, BLACK}, Cell{CASTLE, BLACK}},
	}
}
