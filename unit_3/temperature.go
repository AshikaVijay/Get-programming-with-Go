package main

import (
	"fmt"
)

const (
	line   = "======================="
	row    = "| %5s | %5s |\n"
	number = "%.1f"
)

type getRowFn func(rows int) (string, string) //pg 112 - first class function

func drawTable(celcSymbol string, fahSymbol string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf("%5s", celcSymbol)
	fmt.Printf("%8s \n", fahSymbol)
	fmt.Println(line)
	for r := 0; r < rows; r++ {
		var firstColumn, secondColumn = getRow(r)
		fmt.Printf(row, firstColumn, secondColumn)
	}
	fmt.Println(line)
}

type celsius float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit { //method celsius to fahrenheit
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius { //method fahrenheit to celsius
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func fahrenheitToCelcius(rows int) (string, string) {
	f := fahrenheit(rows*5 - 40) //start from -40 and increase in steps of 5
	c := f.celsius()
	firstColumn := fmt.Sprintf(number, f) //Sprintf to format to a string without printing
	secondColumn := fmt.Sprintf(number, c)
	return firstColumn, secondColumn
}
func celciusToFahrenheit(rows int) (string, string) {
	c := celsius(rows*5 - 40)
	f := c.fahrenheit()
	firstColumn := fmt.Sprintf(number, c)
	secondColumn := fmt.Sprintf(number, f)
	return firstColumn, secondColumn
}

func main() {
	drawTable("°C", "°F", 29, celciusToFahrenheit)
	fmt.Println(line)
	drawTable("°F", "°C", 29, fahrenheitToCelcius)
}
