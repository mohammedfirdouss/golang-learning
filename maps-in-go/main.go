package main

import "fmt"

func main() {
    // Create a map
    myMap := make(map[string]int)

    // Add key-value pairs
    myMap["Alice"] = 25
    myMap["Bob"] = 30

    // Access values
    age, exists := myMap["Bob"]
    if exists {
        fmt.Printf("Bob's age: %d\n", age)
    } else {
        fmt.Println("Key not found")
    }
}
