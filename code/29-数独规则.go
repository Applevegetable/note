package main

import (
	"errors"
	"fmt"
	"os"
)

//固定长宽
const (
	columns = 9
	rows    = 9
	empty   = 0
)

//数独中的方格
type Cell struct {
	digit int8
	fixed bool
}

//数独网格
type Grid [rows][columns]Cell

//声明错误
var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrInRow      = errors.New("digit already present in this row")
	ErrInColunm   = errors.New("digit already present in this colunm")
	ErrInRegion   = errors.New("digit already present in this region")
	ErrFixedDigit = errors.New("initial digits cannot be overwritten")
)

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

//创建一个新的数独网络
func NewSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]
			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}

		}
	}
	return &grid
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns {
		return false
	} else {
		return true
	}
}

func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}
func (g *Grid) inColunm(colunm int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][colunm].digit == digit {
			return true
		}
	}
	return false
}
func (g *Grid) inRegion(row, colunm int, digit int8) bool {
	startRow, startColumn := row/3*3, colunm/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}
func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}

//Set将数字放到数独网格中
func (g *Grid) Set(row, column int, digit int8) error {
	switch {

	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.isFixed(row, column):
		return ErrFixedDigit
	case g.inRow(row, digit):
		return ErrInRow
	case g.inColunm(column, digit):
		return ErrInColunm
	case g.inRegion(row, column, digit):
		return ErrInRegion
	}
	g[row][column].digit = digit
	return nil
}
func (g *Grid) Clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case g.isFixed(row, column):
		return ErrFixedDigit
	}
	g[row][column].digit = empty
	return nil
}

func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	err := s.Set(1, 1, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, row := range s {
		fmt.Println(row)
	}
}
