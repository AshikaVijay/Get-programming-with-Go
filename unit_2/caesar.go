package main

import (
	"fmt"
)

func caesar() {

	message := "L fdph, L vdz, L frqtxhuhg"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c - 3
		} else if c >= 'A' && c <= 'Z' {
			c = c - 3
		}
		fmt.Printf("%c", c)
	}
	fmt.Print("\n")
}
