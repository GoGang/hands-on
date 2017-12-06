package main

import "fmt"

func main() {

	ubc := make(chan string, 2) // buffer de 2 messages : non bloquant
	ubc <- "hello\n"
	ubc <- "world\n"

	fmt.Print(<-ubc)
	fmt.Print(<-ubc)

	bc := make(chan string) // bloquant
	bc <- "hello\n"

	fmt.Print(<-bc)
}
