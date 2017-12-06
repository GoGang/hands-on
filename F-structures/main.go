package main

import "fmt"

type Talk struct {
	Name        string
	Description string
	Properties  map[string]interface{} // dictionnaire : map[type de la clé]type de la valeur
}

func main() {
	talk := Talk{
		Name:        "Golang",
		Description: "A quick tour of golang",
		Properties: map[string]interface{}{
			"foo":   "bar",    // chaine
			"hello": "world!", // chaine
			"num":   1,        // int !
		},
	}

	fmt.Printf("talk: %+v \n", talk)
}
