package main

import "fmt"

func main() {
    // Create a map
    myMap := make(map[string]int)

    // Add key-value pairs
    myMap["Alice"] = 25

    // Access values
    age, exists := myMap["Alice"]
    if exists {
        fmt.Printf("Alice's age: %d\n", age)
    } else {
        fmt.Println("Key not found")
    }

}
