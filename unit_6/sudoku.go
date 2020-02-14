package main

import (
	"errors"
	"fmt"
	"os"
)

// type grid [9][9]int8

const (
	rows      = 9
	columns   = 9
	emptyCell = 0
)

type Cell struct {
	digit int
	fixed bool
}

//a sudoku grid
type sudokuGrid [rows][columns]Cell

//makes a new grid - constructor function
func NewSudoku(array [rows][columns]int) *sudokuGrid {
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

//check if It's a value number
func (grid *sudokuGrid) isValidDigit(digit int) bool {
	//error for words
	//error for non-integars
	if digit >= 1 || digit <= 9 {
		return true
	}
	return false

}

//check if value set for row is between 0-8
func (grid *sudokuGrid) isRowBound(row int) bool {
	if row >= 0 || row < 9 {
		return true
	}
	return false
}

//check if value set for column is between 0-8
func (grid *sudokuGrid) isColumnBound(column int) bool {
	if column >= 0 || column < 9 {
		return true
	}
	return false

}

//check if value is already in the row
func (grid *sudokuGrid) inRow(row, digit int) bool {
	for i := 0; i < columns; i++ {
		if digit == grid[row][i].digit {
			return true
		}
	}
	return false
}

//check if value is already in the column
func (grid *sudokuGrid) inColumn(column, digit int) bool {
	for j := 0; j < rows; j++ {
		if digit == grid[j][column].digit {
			return true
		}
	}
	return false

}

//check if value is in the 3x3 region
func (grid *sudokuGrid) in3x3(row, column, digit int) bool {
	if row >= 0 && row <= 2 {
		row = 0
	} else if row >= 3 && row <= 5 {
		row = 3
	} else if row >= 6 && row <= 8 {
		row = 6
	} else {
		fmt.Print("no row given \n")
	}

	if column >= 0 && column <= 2 {
		column = 0
	} else if column >= 3 && column <= 5 {
		column = 3
	} else if column >= 6 && column <= 8 {
		column = 6
	} else {
		fmt.Print("no column given \n")
	}

	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			if digit == grid[i][j].digit {
				return true
			}
		}
	}

	return false
}

//check if value is fixed from the start
func (grid *sudokuGrid) Fixed(row, column int) bool {
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

func (grid *sudokuGrid) SetaValue(row, column, digit int) error {
	switch {
	case !grid.isValidDigit(digit):
		return ErrDigit
	case !grid.isRowBound(row):
		return ErrRowBounds
	case !grid.isColumnBound(column):
		return ErrColumnBounds
	case grid.inRow(row, digit):
		return ErrRow
	case grid.inColumn(column, digit):
		return ErrColumn
	case grid.in3x3(row, column, digit):
		return Errin3x3
	case grid.Fixed(row, column):
		return ErrFixedDigit
	}

	grid[row][column].digit = digit
	return nil
}

func main() {

	//var array [9][9]int

	sudoku := NewSudoku([rows][columns]int{ //initial digits set
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

	//sudoku.SetaValue(1, 1, 9)

	err := sudoku.SetaValue(1, 1, 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range sudoku {
		fmt.Println(row)
	}

}
