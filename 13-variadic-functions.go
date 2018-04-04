/* Variadic functions can be called with any number of trailing arguments. */

package main

import "fmt"

func sum(numbers ...int) int {
    fmt.Println("Numbers to be summed.")
    fmt.Println(numbers, " ")

    var total int

    for _, number := range numbers {
        total += number
    }
    fmt.Println("Returning summation.")
    return total
}

func main() {

    result := sum(1, 2)
    fmt.Println("Sum of 1 and 2 = ", result)

    result = sum(1, 2, 3, 4)
    fmt.Println("Sum of 1, 2, 3 and 4 = ", result)

    numbers := []int {1, 2, 3, 4, 5, 6}
    result = sum(numbers ...)
    fmt.Println("Sum of 1, 2, 3, 4, 5 and 6 = ", result)

}