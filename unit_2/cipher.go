package main

import (
	"fmt"
	"strings"
)

// func shift(c byte) {
// 	return (c - 'A') % 26
// }

func cipher() {

	plainText := "your message goes here"
	keyword := "GOLANG"
	j := 0
	cipherText := ""

	replacePlainText := strings.Replace(plainText, " ", "", -1) //gets rid of spaces
	plainText = strings.ToUpper(replacePlainText)               //all upper case just like the keyword
	replaceKeyword := strings.Replace(keyword, " ", "", -1)
	keyword = strings.ToUpper(replaceKeyword)

	for i := 0; i < len(plainText); i++ {
		c := plainText[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 'A'
			kj := keyword[j] - 'A'
			c = (c+kj)%26 + 'A' //cipher letter + kj letter
			j++
			j %= len(keyword)
		}
		cipherText = cipherText + string(c)
	}
	fmt.Println(cipherText)
}
