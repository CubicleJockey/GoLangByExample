package main

import (
    "math"
    "fmt"
    "reflect"
)

type geometry interface {
    area() float64
    perimeter() float64
}

//Rectangle START

//previously defined in 18-methods as int, hence rename
type rectangle64 struct {
    width float64
    height float64
}

func (r rectangle64) area() float64 {
    return r.width * r.height
}

func (r rectangle64) perimeter() float64 {
    return (2 * r.width) + (2 * r.height)
}

//Rectangle END

//Circle START
type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
//Circle END

func measure(g geometry) {
    fmt.Println("Type: ", reflect.TypeOf(g))
    fmt.Println("Type Value: ", g)
    fmt.Println("Area: ", g.area())
    fmt.Println("Perimeter: ", g.perimeter())
    fmt.Println()
}

func main(){
    rect := rectangle64{3, 4}
    measure(rect)

    cir := circle{5}
    measure(cir)
}




