/*
Go supports methods defined on struct types.

Methods can be defined for either pointer or value receiver types.

Go automatically handles conversion between values and pointers for method calls.
You may want to use a pointer receiver type to avoid copying on method calls or to
allow the method to mutate the receiving struct.
*/

package main

import "fmt"

type rectangle struct {
    width int
    height int
}

func (rect *rectangle) area() int {
    return rect.width * rect.height
}

func (rect rectangle) perimeter() int {
    return (2 * rect.width) + (2 * rect.height)
}

func main() {

    rect := rectangle{10, 5}

    fmt.Println("Area: ", rect.area())
    fmt.Println("Perimeter: ", rect.perimeter())

    rectanglePointer := &rect

    fmt.Println("Pointer Area: ", rectanglePointer.area())
    fmt.Println("Pointer Perimeter: ", rectanglePointer.perimeter())
}