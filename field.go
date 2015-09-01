package main

import (
	"fmt"
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
	unit  Unit
	color Color
}

func (c Cell) String() string { return fmt.Sprint("[", c.unit, c.color, "],") }

//Поле
type Field [8][8]Cell

func InitGameField() *Field {
	return &Field{
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
