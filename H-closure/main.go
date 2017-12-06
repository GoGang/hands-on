package main

import "fmt"

func main() {
	// déclaration d'une fonction dans la fonction main
	count := func(name string) int {
		return len(name)
	}
	// func count() int { return len(name) } # ne compile pas.

	fmt.Printf("Count: %d\n", count("Golang"))

	hello := "Hello"

	join := func(name string) string {
		return fmt.Sprintf("%s %s", hello, name)
	}

	fmt.Printf("Join: %s\n", join("Golang"))
}
