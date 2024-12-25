// main.go
package main

import "fmt"

// Define a struct to represent a person
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Create an instance of the struct
    person := Person{
        Name: "John Doe",
        Age:  30,
        City: "New York",
    }

    // Access and print struct fields
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    fmt.Println("City:", person.City)
}
