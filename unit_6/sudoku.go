package main

import (
	"errors"
	"fmt"
)

// type grid [9][9]int8

const (
	rows      = 9
	columns   = 9
	emptyCell = 0
)

type Cell struct {
	digit int8
	fixed bool
}

//a sudoku grid
type sudokuGrid [rows][columns]Cell

//makes a new grid - constructor function
func NewSudoku(array [rows][columns]int8) *sudokuGrid {
	var grid sudokuGrid
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			if array[i][j] > 0 {
				grid[i][j].fixed = true
				grid[i][j].digit = array[i][j]
			} else {
				grid[i][j].fixed = false
			}
		}
	}
	return &grid
}

//isrowbound method
//is columnbound method
//method to check if in horizontal row
//method to check if in vertical row
//method to check in 3x3 row
//method to cannot fix digit
//set digit in a specific location (row, col, digit)
//clear a digit from a location

// Set a digit on a Sudoku grid.

//is it a fixed number
func (grid *sudokuGrid) isValidDigit(digit int8) bool {
	if digit <= 1 || digit >= 9 {
		return true
	}
	return false
}

func (grid *sudokuGrid) isRowBound(row int8) bool {
	if row >= 1 || row <= 9 {
		return true
	}
	return false
}
func (grid *sudokuGrid) isColumnBound(column int8) bool {
	if column >= 1 || column <= 9 {
		return true
	}
	return false

}
func (grid *sudokuGrid) inRow(row, digit int8) bool {
	for i := 0; i < columns; i++ {
		if digit == grid[row][i].digit {
			return true
		}
	}
	return false
}

func (grid *sudokuGrid) inColumn(column, digit int8) bool {
	for j := 0; j < rows; j++ {
		if digit == grid[j][column].digit {
			return true
		}
	}
	return false

}

func (grid *sudokuGrid) in3x3(row, column, digit int8) bool {
	firstRow := row / 9
	firstColumn := column / 9

	for i := firstRow; i < firstRow+3; i++ {
		for j := firstColumn; j < firstColumn+3; j++ {
			if digit == grid[i][j].digit {
				return true
			}
		}
	}
	return false
}

func (grid *sudokuGrid) Fixed(row, column int8) bool {
	if grid[row][column].fixed {
		return true
	}
	return false
}

//ERRORS
var (
	ErrDigit        = errors.New("Not a valid digit")
	ErrRowBounds    = errors.New("out of row bounds")
	ErrColumnBounds = errors.New("out of column bounds")
	ErrRow          = errors.New("digit already in row")
	ErrColumn       = errors.New("digit already in column")
	Errin3x3        = errors.New("digit already present in this 3x3 region")
	ErrFixedDigit   = errors.New("cannot overwrite a fixed digit")
)

func (grid *sudokuGrid) SetaValue(row, column, digit int8) error {
	switch {
	case !grid.isValidDigit(digit):
		return ErrDigit
	case !grid.isRowBound(row):
		return ErrRowBounds
	case !grid.isColumnBound(column):
		return ErrColumnBounds
	case !grid.inRow(row, digit):
		return ErrRow
	case !grid.inColumn(column, digit):
		return ErrColumn
	case !grid.in3x3(row, column, digit):
		return Errin3x3
	case !grid.Fixed(row, column):
		return ErrFixedDigit
	}

	grid[row][column].digit = digit
	return nil
}

func main() {

	//ar array [9][9]int

	sudoku := NewSudoku([rows][columns]int8{ //initial digits set
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

	SetaValue(1, 1, 1)

	fmt.Print(sudoku)

}
