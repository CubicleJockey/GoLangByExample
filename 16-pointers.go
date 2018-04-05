/* Go supports pointers, allowing you to pass references to values and records within your program. */

package main

import "fmt"

func zeroValue(value int) {
    value = 0
}

func zeroPointer(value *int) {
    *value = 0
}

func main() {

    i := 1

    fmt.Println("initial: ", i)

    zeroValue(i)
    fmt.Println("zeroValue called with result: ", i)

    zeroPointer(&i)
    fmt.Println("zeroPointer called with result: ", i)

    fmt.Println("pointer memory address: ", &i)

}