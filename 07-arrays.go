/*
In Go, an array is a numbered sequence of elements of a specific length.


Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.

We can set a value at an index using the array[index] = value syntax, and get a value with array[index].

The builtin len returns the length of an array.

Use this syntax to declare and initialize an array in one line.

Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.

*/

package main

import "fmt"

func main() {
	
	var numbers [5] int
	fmt.Println("Initialized", numbers)
	
	const val = 100
	fmt.Println("Set last item to: ", val)
	numbers[4] = val
	
	fmt.Println("Get last item: ", numbers[4]) 
	fmt.Println("Get entire array: ", numbers)
		
	fmt.Println("\n")
	
	moreNumbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Declared and Initialized at same time: ", moreNumbers)
	
	var twoDim [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDim[i][j] = i + j
		}
	}
	
	fmt.Println("2-dim: ", twoDim)
}