package main

import (
	"fmt"
	"math/rand"
)

func piggy2() {

	piggyBank := 0
	nickels := 5
	dimes := 10
	quarters := 25 //cents

	for piggyBank < 2000 {
		switch rand.Intn(3) {
		case 0:
			piggyBank += nickels
		case 1:
			piggyBank += dimes
		case 2:
			piggyBank += quarters
		}

		piggyBankDollars := piggyBank / 100
		piggyBankCents := piggyBank % 100

		fmt.Printf("$%d.%02d\n", piggyBankDollars, piggyBankCents)
	}
}
