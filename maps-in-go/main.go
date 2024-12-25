package main

import "fmt"

func main() {
    // Create a map
    myMap := make(map[string]int)

    // Add key-value pairs
    myMap["Alice"] = 25
    myMap["Bob"] = 30
	myMap["Firdous"] = 50

    // Access values
    age, exists := myMap["Alice"]
    if exists {
        fmt.Printf("Alice's age: %d\n", age)
    } else {
        fmt.Println("Key not found")
    }

    // Iterate over the map
    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }

    // Delete a key-value pair
    delete(myMap, "Bob")
	delete(myMap, "Firdous")
    fmt.Println("After deletion:", myMap)
}
