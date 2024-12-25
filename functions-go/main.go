package main

import "fmt"

// A function with no inputs and no outputs
func greet() {
    fmt.Println("Hello, World!")
}

// A function with inputs and a single output
func add(a int, b int) int {
    return a + b
}

// A function with inputs and multiple outputs
func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    greet() // Call greet

    result := add(3, 4) // Call add
    fmt.Println("3 + 4 =", result)

    quotient, err := divide(10, 2) // Call divide
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 2 =", quotient)
    }
}