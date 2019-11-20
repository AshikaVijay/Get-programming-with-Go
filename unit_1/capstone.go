package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var distance = 62100000 //distance in km to mars
	var company = ""
	var trip = ""

	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")

	for i := 0; i < 10; i++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		var speed = rand.Intn(15) + 16      //speed in km/s chosen between 16-30km/s
		var days = distance / speed / 86400 //convert seconds to days
		var price = speed + 20              //faster trips more expensive

		if rand.Intn(2) == 0 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}

		fmt.Printf("%-16v %4v %-10v $%4v \n", company, days, trip, price)
	}
}
