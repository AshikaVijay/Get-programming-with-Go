package main

import (
	"fmt"
	"math/rand"
)

func piggy() {

	piggyBank := 0.0
	nickels := 0.05
	dimes := 0.10
	quarters := 0.25

	for piggyBank < 20.00 {
		switch rand.Intn(3) {
		case 0:
			piggyBank += nickels
		case 1:
			piggyBank += dimes
		case 2:
			piggyBank += quarters
		}

		fmt.Printf("$%5.2f\n", piggyBank)
	}

}
