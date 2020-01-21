package main

import "fmt"

type Planets []string //Planets attaches method to []string

func (planets Planets) terraform() { //Planets type with a terraform method. //planets is receiver of terraform and Planets is receiver type
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	Planets(planets[3:4]).terraform() //converts the planets to the 'Planets' type, then calls 'terrafrom' method.
	Planets(planets[6:]).terraform()

	fmt.Println(planets)
}
