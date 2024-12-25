package main

	import "fmt"
	
	func main() {
		// Declaring an array
		var arr [5]int
	
		// Assign values to the array
		arr[0] = 10
		arr[1] = 20
		arr[2] = 30
		arr[3] = 40
		arr[4] = 50
	
		// Iterating over the array
		for i, val := range arr {
			fmt.Printf("Index %d: Value %d\n", i, val)
		}
	}
