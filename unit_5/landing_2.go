package main

import (
	"fmt"
	"os"
)

//exitOnError prints any errors and exits.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {

	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28,
		21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354,
		28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137,
		26, 30.12, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54,
		0, 'E'})
	fmt.Println("Spirit", spirit)
	fmt.Println("Opportunity", opportunity)
	fmt.Println("Curiosity", curiosity)
	fmt.Println("InSight", insight)
}
