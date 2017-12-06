package main

import (
	"fmt"
)

func main() {

	c := make(chan string) // on alloue et on affecte un canal de chaines de caractères

	alim := func(c chan string) { // on alimente le canal dans une goroutine
		c <- "hello\n"
		c <- "world\n"
	}

	go alim(c)

	for {
		fmt.Print(<-c) // on va piocher dedans
	}

	// le programme s'arrête correctement ?
}
