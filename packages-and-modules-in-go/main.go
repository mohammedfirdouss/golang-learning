package main

import (
	"fmt"
	"example.com/mymodule/utility" 
)

func main() {
	fmt.Println("Using a package and module in Go!")
	result := utility.Add(5, 3)
	fmt.Println("The sum is:", result)
}
