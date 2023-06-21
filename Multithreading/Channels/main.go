package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) //Vazio

	go func() {
		canal <- "Olá Mundo" //Canal cheio
	}()

	msg := <-canal // Canal vazio
	fmt.Println(msg)
}
