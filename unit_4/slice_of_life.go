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

func (u Universe) Neighbours(w, h int) int { //counts the number of adjacent cells that are alive

	var N, S, E, W, NE, SE, SW, NW bool

	N = u.Alive(w, h+1)
	S = u.Alive(w, h-1)
	E = u.Alive(w+1, h)
	W = u.Alive(w-1, h)
	NE = u.Alive(w+1, h+1)
	SE = u.Alive(w+1, h-1)
	SW = u.Alive(w-1, h-1)
	NW = u.Alive(w-1, h+1)

	list := []bool{N, S, E, W, NE, NW, SE, SW}

	total := 0
	for i := 0; i < len(list); i++ {
		total = total + Btoi(list[i])
	}

	return total

}

func (u Universe) Next(w, h int) bool { //whether cell is alive or not in next generation

	total := u.Neighbours(w, h)
	return (total == 2 || total == 3) && u.Alive(w, h)

}

func main() {

	u := NewUniverse()
	u.Seed()
	//u.Show()
	u.Alive(56, 13)
	u.Neighbours(56, 13)
	u.Next(30, 14)

}
