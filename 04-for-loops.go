/*

for is Go’s only looping construct. Here are three basic types of for loops.

The most basic type, with a single condition.

A classic initial/condition/after for loop.

for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

You can also continue to the next iteration of the loop.

*/

package main

import "fmt"

func main(){

	fmt.Println("Basic for loop single condition")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	
	fmt.Println("Classic for loop")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	
	fmt.Println("For without condition")
	for {
		fmt.Println("Starting loop...")
		break;
		fmt.Println("break was called loop ended.")
	}
	
	fmt.Println("Continue to next loop iteration")
	for n := 0; n <=5; n++ {
		if n % 2 == 0 { 
			fmt.Println("continue invoked skip rest of loop")
			continue
		}
		fmt.Println(n)
	}
}