package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var number = 33 //number to guess

	for {
		var num = rand.Intn(100) + 1
		if num > number {
			fmt.Printf("%v is too big \n", num)
		} else if num < number {
			fmt.Printf("%v is too small \n", num)
		} else {
			fmt.Print("Numbers match \n")
			break
		}
	}
}
