package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers3(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	wg.Done()
}

func printLetters3(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers3(&wg)
	go printLetters3(&wg)
	wg.Wait()
}
