package main
	import "fmt"


	func main() {
		// Create a slice using make
		slice := make([]int, 3)
		slice[0] = 10
		slice[1] = 20
		slice[2] = 30
	
		// Slice an existing array
		arr := [5]int{1, 2, 3, 4, 5}
		mySlice := arr[1:4]
		fmt.Println("mySlice:", mySlice)
	}
	
