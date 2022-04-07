package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello"
}

func callerB(c chan string) {
	c <- "Hola"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	for i := 0; i < 5; i++ {
		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("Default")
		}
		time.Sleep(1 * time.Microsecond)
	}
}
