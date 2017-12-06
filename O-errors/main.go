package main

import (
	"errors"
	"fmt"
)

// on créé une erreur maison
var InvalidSumError = errors.New("Invalid Sum, only positive values are required")

func PositiveSum(i, y int) (int, error) {
	defer func() {
		fmt.Printf("Call after function return\n")
	}()

	if i < 0 || y < 0 {
		return 0, InvalidSumError
	}

	return i + y, nil
}

func main() {
	if value, err := PositiveSum(1, -1); InvalidSumError == err { // on doit lever l'erreur
		fmt.Printf("Error: %s\n", err.Error()) // on affiche le message
	} else {
		fmt.Printf("Value: %+v\n", value) // sinon on affiche la valeur (notez la directive de formatage)
	}

	value, err := PositiveSum(1, 1)

	fmt.Printf("Value: %+v, Error: %s\n", value, err)
}
