package main

import (
	"fmt"
)

func canis() {

	const distance = 2.36e17
	const lightYear = 9.45e12

	const convert = distance / lightYear

	fmt.Printf("%v\n", convert)

}
