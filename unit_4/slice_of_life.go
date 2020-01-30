package main

import (
	"fmt"
	"math/rand"
	"time"
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
	for h := 1; h < height; h++ {
		for w := 1; w < width; w++ {
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

//clear grid
func (u Universe) Show() {
	fmt.Print("\x0c", u.ToString())
}

//seed 25%
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

//determines whether a cell is dead or alive
func (u Universe) Alive(w, h int) bool {
	if h >= 0 || h < 15 {
		h = h % height
	}

	if w >= 0 || w < 80 {
		w = w % width
	}

	return u[h][w]
}

//Change boolean to integer
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

//counts the number of adjacent cells that are alive
func (u Universe) Neighbours(w, h int) int {

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

//whether cell is alive or not in next generation
func (u Universe) Next(w, h int) bool {

	total := u.Neighbours(w, h)
	return (total == 2 || total == 3) && u.Alive(w, h)

}

func Step(a, b Universe) {
	for h := 1; h < height; h++ {
		for w := 1; w < width; w++ {
			b[h][w] = a.Next(w, h)
		}
	}
}

func main() {

	a := NewUniverse()
	b := NewUniverse()

	a.Seed()

	for i := 1; i < 10; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(1 * time.Second)
		a, b = b, a //Swap
	}

}
