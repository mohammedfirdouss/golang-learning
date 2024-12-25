package main

import "fmt"

// Define an interface
type Greeter interface {
    Greet() string
}

// Implement the interface with a struct
type Person struct {
    Name string
}

func (p Person) Greet() string {
    return "Hello, " + p.Name
}


func main() {
    var greeter Greeter = Person{Name: "Alice"}
    fmt.Println(greeter.Greet())
}