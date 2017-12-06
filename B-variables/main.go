package main

import "fmt"

// constantes
const w = 3

// variables globales au package
var x = 2

func main() {
	// variables locales
	var y int64 // valeur par défaut (0)
	z := 1      // allocation + assignation via inférence de type (avec les ":" )
	z = 3       // réassignation d'une variable déjà allouée
	x = 3
	// w = 3 # échoue sans les ":"

	fmt.Printf("w: %d, type: %T\n", w, w) // formatage : %d entier, %T type
	fmt.Printf("x: %d, type: %T\n", x, x)
	fmt.Printf("y: %d, type: %T\n", y, y)
	fmt.Printf("z: %d, type: %T\n", z, z)
}
