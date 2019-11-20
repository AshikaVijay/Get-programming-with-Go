package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {

	for i := 0; i < 100; i++ {
		year := 2019 + rand.Intn(100)
		leapYear := year%400 == 0 || (year%100 != 0 && year%4 == 0)
		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2:
			if leapYear {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}

		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}
