package main

import (
	"fmt"
	"math/rand"
	"time"
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
	switch rand.Intn(3) {
	case 0:
		return "wotsits"
	case 1:
		return "biscuits"
	default:
		return "chocolate"
	}
}

func (p Pig) move() string {
	switch rand.Intn(2) {
	case 0:
		return "taps"
	case 1:
		return "springs"
	default:
		return "slides"
	}
}

//dinosaur name, move, eat
type Dinosaur struct {
	name string
}

func (d Dinosaur) String() string {
	return d.name
}

func (d Dinosaur) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "gummy bears"
	case 1:
		return "noodles"
	default:
		return "strawberry"
	}
}

func (d Dinosaur) move() string {
	switch rand.Intn(2) {
	case 0:
		return "stamps"
	case 1:
		return "shuffles"
	default:
		return "steps"
	}
}

type animal interface {
	eat() string
	move() string
}

//
func choose(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v eats %v \n", a, a.eat())
	case 1:
		fmt.Printf("%v %v \n", a, a.move())
	}
}

func main() {

	animals := []animal{
		Pig{name: "Hamm"},
		Dinosaur{name: "Rex"},
	}

	var sol int
	var hour int

	const sunrise = 6
	const sunset = 18

	fmt.Print("Sol 1! \n")

	for {

		//prints every hour
		fmt.Printf("At %2d:00 ", hour)

		if hour >= sunset || hour <= sunrise {
			fmt.Println("zzzzzzzz")
		} else {
			i := rand.Intn(len(animals))
			choose(animals[i])

		}

		time.Sleep(400 * time.Millisecond)

		//3, 24 hour sols
		hour = hour + 1
		if hour >= 24 {
			hour = 0
			if sol == 0 {
				fmt.Print("Sol 2! \n")
			} else if sol == 1 {
				fmt.Print("Sol 3! \n")
			}
			sol++
			if sol >= 3 {
				break
			}
		}
	}

}
