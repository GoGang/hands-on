package main

import "fmt"

func main() {

	i := 0
	for { // équivalent d'un while(true). Pas de while en Go.
		if i > 9 {
			fmt.Printf("i: %d, type: %T\n", i, i)
			break // on sort par là
		}
		i++
	}

	for i := 0; i < 20; i++ { // le i est local
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\ni: %d \n", i)

	for ; i < 20; i++ { // utilisation du i global
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\ni: %d \n", i)
}
