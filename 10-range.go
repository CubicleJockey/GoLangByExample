/*
    range iterates over elements in a variety of data structures.
*/

package main

import "fmt"

func main(){

    numbers := []int{2, 3, 4}
    var sum int
    for index, number := range numbers {
        fmt.Println("Index [", index, "] has value: ", number)
        sum += number
    }
}

