package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // Adding two goroutines to the wait group
	
	go printNumbers(&wg, "Goroutine 1")
	go printNumbers(&wg, "Goroutine 2")

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines completed.")
}