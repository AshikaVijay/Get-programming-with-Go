package main

import (
	"fmt"
)

func Source(downstream chan string) {
	array := []string{"hello world", "hello world", "hi world"}
	for i := range array {
		downstream <- array[i] //sends array element downstream
	}
	close(downstream) //indicate when done
}

func NoRepeatFilter(upstream, downstream chan string) {
	item := ""
	for newItem := range upstream {
		if newItem != item {
			downstream <- newItem
			item = newItem
		}
	}
	close(downstream)
}

func Print(upstream chan string) {
	for newItem := range upstream {
		fmt.Println(newItem)
	}
}

func main() {
	c0 := make(chan string) //dowstream
	c1 := make(chan string) //upstream
	go Source(c0)
	go NoRepeatFilter(c0, c1)
	Print(c1)
}
