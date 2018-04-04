package main

import "fmt"

func values() (string, string) {
    return "First Return Value", "Second Return Value"
}

func main() {

    first, second := values()
    fmt.Println(first)
    fmt.Println(second)

    // ignore first value
    // If you only want a subset of the returned values, use the blank identifier _.
    _, collected := values()
    fmt.Println(collected)
}
