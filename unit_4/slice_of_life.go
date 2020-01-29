package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 80
	height = 15
)

//Universe 2d field of cells
type Universe [][]bool

//NewUniverse fn returns an empty universe
func NewUniverse() Universe {
	u := make(Universe, height) //15
	for i := range u {
		u[i] = make([]bool, width) //80

	}
	return u
}

//change ints to strings to print * and ''
func (u Universe) ToString() string {
	var cell rune
	NewArray := make([]rune, width*height)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if u[h][w] == true {
				cell = '*'
			} else {
				cell = ' '
			}
			NewArray = append(NewArray, cell)
		}
		NewArray = append(NewArray, '\n')
	}
	return string(NewArray)
}

//show grid
func (u Universe) Show() {
	fmt.Print(u.ToString())
}

//seed 25%
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

func (u Universe) Alive(w, h int) bool { //determines whether a cell is dead or alive
	if h <= 0 || h > 15 {
		fmt.Print("Please put a postive integer less than 50 \n")
	} else {
		h = h % height
	}

	if w <= 0 || w > 80 {
		fmt.Print("Please put a postive integer less than 80 \n")
	} else {
		w = w % width
	}
	return u[h][w]
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (u Universe) Neighbours(w, h int) int {
	//counts the number of adjacent cells that are alive
	var north, south, east, west, NorthEast, SouthEast, SouthWest, NorthWest bool

	north = u.Alive(w, h+1)
	south = u.Alive(w, h-1)
	east = u.Alive(w+1, h)
	west = u.Alive(w-1, h)
	NorthEast = u.Alive(w+1, h+1)
	SouthEast = u.Alive(w+1, h-1)
	SouthWest = u.Alive(w-1, h-1)
	NorthWest = u.Alive(w-1, h+1)

	list := []bool{north, south, east, west, NorthEast, NorthWest, SouthEast, SouthWest}

	total := 0
	for i := 0; i < len(list); i++ {
		total = total + Btoi(list[i])
	}

	return total
}

func (u Universe) Next(w, h int) bool {

if 

//live cell with less than 2 live neighbours dies
//a  live cell with 2 or 3 live neighbours lives on to the next generation 
// a live cell with >3 neighbours dies
//a dead cell with =3 neighbours becomes a live cell

return next 
}

func main() {

	u := NewUniverse()
	u.Seed()
	//u.Show()
	//u.Alive(10, 4)
	u.Neighbours(30, 14)

}
