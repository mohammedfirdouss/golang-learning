# Slices in Go

Slices in Go are dynamic and more flexible than arrays. They represent a portion of an underlying array.

- To create a slice, you can make use of the *make* function.
- If the underlying array is not large enough, you can make use of the *append* function
- range is used with slices  to iterate over their elements. When used with slices, it provides the index and value of each element.



## My Points:
1. Declare slices using `[]Type`.
2. Slices are dynamically sized and can grow or shrink.
3. Use `make` to create slices, or slice an existing array.
