package main

import (
	"fmt"
	"math"
)

//CONVERT TO DECIMAL
type location struct {
	lat, long float64
}

type coordinate struct {
	d, m, s float64
	h       rune
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

//FIND DISTANCES
type world struct {
	radius float64
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

var mars = world{radius: 3389.5}
var earth = world{radius: 6371}

func main() {

	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28,
		21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354,
		28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137,
		26, 30.12, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54,
		0, 'E'})

	fmt.Printf("Spirit to Opportunity: %.2f km\n", mars.distance(spirit, opportunity))
	fmt.Printf("Spirit to Curiosity: %.2f km\n", mars.distance(spirit, curiosity))
	fmt.Printf("Spirit to Insight: %.2f km\n", mars.distance(spirit, insight))
	fmt.Printf("Opportunity to Curiosity: %.2f km\n", mars.distance(opportunity, curiosity))
	fmt.Printf("Opportunity to Insight: %.2f km\n", mars.distance(opportunity, insight))
	fmt.Printf("Curiosity to Insight: %.2f km\n", mars.distance(curiosity, insight))

	london := newLocation(coordinate{51, 30, 0.0, 'N'}, coordinate{0, 8,
		0, 'W'})
	paris := newLocation(coordinate{48, 51, 0.0, 'N'}, coordinate{2, 21,
		0, 'E'})

	fmt.Printf("London to Paris: %.2f km\n", earth.distance(london, paris))

	manchester := newLocation(coordinate{53, 29, 2.2524, 'N'}, coordinate{2, 14,
		40.7184, 'W'})

	fmt.Printf("Manchester to London: %.2f km\n", earth.distance(manchester, london))

	mountSharp := newLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51,
		0, 'E'})
	OlympusMons := newLocation(coordinate{18, 39, 0.0, 'N'}, coordinate{226, 12,
		0, 'E'})

	fmt.Printf("Mount Sharp to Olympus Mons: %.2f km\n", mars.distance(mountSharp, OlympusMons))

}
