package main

import (
	"fmt"
	"strings"
)

func SourceGo(downstream chan string) {

	for _, s := range []string{"hello hi world"} {
		downstream <- s
	}
	close(downstream) //indicate when done

}

func splitString(upstream, downstream chan string) {
	for up := range upstream { //get all upstream
		split := strings.Fields(up) //get ["hello" "hi" "world"]
		for _, s := range split {   //splits this up
			downstream <- s //sends downstream
		}
	}
	close(downstream) //indicate when done
}

func PrintGo(upstream chan string) {
	for words := range upstream {
		fmt.Println(words)
	}
}

//close(upstream)

func main() {
	c0 := make(chan string) //dowstream
	c1 := make(chan string) //upstream
	go SourceGo(c0)
	go splitString(c0, c1)
	PrintGo(c1)
}
