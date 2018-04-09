/*

Go’s sort package implements sorting for builtins and user-defined types.
We’ll look at sorting for builtins first.
*/

package main

import (
    "sort"
    "fmt"
)

func main() {

    characters := []string { "c", "a", "b" }

    fmt.Println("Original Characters:", characters)
    fmt.Println("Sorting...")

    sort.Strings(characters)

    fmt.Println("Sorted: ", sort.StringsAreSorted(characters))
    fmt.Println("Sorted Characters", characters)

    integers := []int  {3, 4, 2, 1, 5 }

    fmt.Println("Original Integers:", integers)
    fmt.Println("Sorting...")

    sort.Ints(integers)

    fmt.Println("Sorted:", sort.IntsAreSorted(integers))
    fmt.Println("Sorted integers:", integers)

}