package main

import (
    "errors"
    "fmt"
    "strconv"
)

func DoWork(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New(strconv.Itoa(arg) + " is not a valid value")
    }
    return arg + 3, nil
}


func main() {

    for _, number := range []int {1, 7, 42} {
        if num, err := DoWork(number); err != nil {
            fmt.Println("DoWork input: ", number)
            fmt.Println("DoWork result: ", err)
            fmt.Println()
        } else {
            fmt.Println("DoWork input: ", number)
            fmt.Println("DoWork result: ", num)
            fmt.Println()
        }
    }
}
