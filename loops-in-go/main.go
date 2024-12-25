package main

import "fmt"

func main() {
    // Example Basic for loop
    for i := 0; i < 5; i++ {
        fmt.Println("Basic loop iteration:", i)
    }

    // Infinite loop
    i := 0
    for {
        if i%2 == 0 {
            i++
            continue // Skip printing for even numbers
        }
        fmt.Println("Infinite loop iteration (odd):", i)
        i++
        if i >= 5 {
            break // stops the current iteration and exits the loop
        }
    }
}