package main

import "fmt"

//Addition
func plus(lhs int, rhs int) int {
    return lhs + rhs
}

//Double Addition
func plusPlus(lhs, mid, rhs int) int {
    return lhs + mid + rhs
}

func main(){

    result := plus(2, 2)
    fmt.Println("2 + 2 = ", result)

    result = plusPlus(1, 2, 3)
    fmt.Println("1 + 2 + 3 = ", result)

}