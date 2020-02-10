package main

import "fmt"

type turtle struct {
	x int
	y int
}

func (t *turtle) up() {
	t.y--
}
func (t *turtle) down() { //positive values go down
	t.y++
}
func (t *turtle) left() {
	t.x--
}
func (t *turtle) right() { //pos values go right
	t.x++
}
func main() {

	var t turtle
	t.up()
	t.left()
	t.left()
	fmt.Println(t)

}
