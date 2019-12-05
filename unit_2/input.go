package main

import (
	"fmt"
)

func input() {

	yesOrNo := "no"

	var answer bool

	switch yesOrNo {
	case "true", "yes", "1":
		answer = true
	case "false", "no", "0":
		answer = false
	default:
		fmt.Println("Error")
	}

	fmt.Println("Answer: ", answer)

}
