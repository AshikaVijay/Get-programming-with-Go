package main

const (
	row    = 9
	column = 9
)

type Celll struct {
	digit int8
	fixed bool
}

//a sudoku grid
type sudokuGridd [rows][columns]Cell

func (grid *sudokuGridd) in3x3(r, c, digit int8) bool {
	firstRow := r / 9
	firstColumn := c / 9

	for i := firstRow; i < firstRow+3; i++ {
		for j := firstColumn; j < firstColumn+3; j++ {
			if digit == grid[i][j].digit {
				return true
			}
		}
	}
	return false
}

func main() {

}
