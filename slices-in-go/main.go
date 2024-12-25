package main
	import "fmt"


	func main() {
		// Create a slice using make
		slice := make([]int, 3)
		slice[0] = 10
		slice[1] = 20
		slice[2] = 30

    	// Append to the slice
    	slice = append(slice, 40, 50)
	
		// Iterate over the slice
		for i, val := range slice {
			fmt.Printf("Index %d: Value %d\n", i, val)
		}
		// Slice an existing array
		arr := [5]int{1, 2, 3, 4, 5}
		mySlice := arr[1:4]
		fmt.Println("mySlice:", mySlice)
	}
	
