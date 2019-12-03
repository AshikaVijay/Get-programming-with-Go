package main

import (
	"fmt"
)

func malacandra() {

	const distance = 56000000
	var time = 28 * 24 //convert days to hours
	var speed = distance / time

	fmt.Println(speed, "km/hr")

}
