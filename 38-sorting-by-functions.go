/*
Sometimes we’ll want to sort a collection by something other than its natural order.

For example, suppose we wanted to sort strings by their length instead of alphabetically.
Here’s an example of custom sorts in Go.

*/

package main

import (
    "fmt"
    "sort"
)

//Define Methods START

type byLength []string

func (str byLength) Len() int {
    return len(str)
}

func (str byLength) Swap(lhs, rhs int) {
    str[lhs], str[rhs] = str[rhs], str[lhs]
}

func (str byLength) Less(lhs, rhs int) bool {
    return len(str[lhs]) < len(str[rhs])
}

//Define Methods END

func main() {

    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}