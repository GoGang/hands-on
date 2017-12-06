package main

import (
	"fmt"
	"time"
)

func main() {

	max := 2
	c := make(chan string, max)

	for _, v := range []int{1, 2, 3, 4, 5} {
		go func(v int) { // chaque goroutine envoie un message et est garbage collectée
			c <- fmt.Sprintf("Message #%d", v)
		}(v)
	}
	go func() { // une seule goroutine récupère tous les messages
		for {
			value := <-c
			fmt.Println(value)
		}
	}()
	time.Sleep(3 * 1000)
}
