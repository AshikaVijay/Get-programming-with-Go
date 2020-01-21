package main

import (
	"fmt"
)

func decipher() {

	cipherText := "POYBXXCIAMWHDMAUMCGSAQNUWLSBSGZ"
	keyword := "MALINKO"
	j := 0
	plainText := ""

	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 'A'
			kj := keyword[j] - 'A'
			c = (c-kj+26)%26 + 'A'
			j++
			j %= len(keyword)
		}
		plainText = plainText + string(c)
	}
	fmt.Println(plainText)
}
