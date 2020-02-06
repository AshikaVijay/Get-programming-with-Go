package main

import (
	"fmt"
	"math/rand"
)

//pig name, move and eat
type Pig struct {
	name string
}

//stringer interface to return name
func (p Pig) String() string {
	return p.name
}

//eat method
func (p Pig) eat() string {
	switch rand.Int(3) {
	case 0:
		return "wotsits"
	case 1:
		return "biscuits"
	case 2:
		return "crisps"
	}
}

func (p Pig) move() string {
	return "shuffle"
}

//dinosaur
type Dinosaur struct {
	name string
}

func (d Dinosaur) String() string {
	return d.name
}

func (d Dinosaur) eat() string {
	switch rand.Int(2) {
	case 0:
		return "gummy bears"
	case 1:
		return "skittles"
	}
}

func (d Dinosaur) move() string {
	return "stamp"
}

type animal interface {
	eat() string
	move() string
}

func choose(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v eats %v \n", a, a.eat())
	case 1:
		fmt.Printf("%v %v \n", a, a.move())
	}
}

const sunrise = 6
const sunset = 18

func main() {
	animals := []animal{
		Pig{name: "Hamm"},
		Dinosaur{name: "Rex"},
	}

	var sol int
	var hour int

}
