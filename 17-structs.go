/*
Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
*/

package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    fmt.Println(person{ "James Howlett", 225 })
    fmt.Println(person{name: "Wayde Wilson", age: 42})
    fmt.Println(person{name: "Jason Voorhees"})

    // An & prefix yields a pointer to the struct
    fmt.Println(&person{name: "Logan", age: 225})

    spiderman := person{"Spider-Man", 28 }
    fmt.Println(spiderman)

    spidermanPointer := &spiderman
    fmt.Println(spiderman)
    fmt.Println(spidermanPointer)

    fmt.Println("Updating the pointer object only...")
    spidermanPointer.age = 42
    fmt.Println("Pointer updated.")

    fmt.Println(spiderman)
    fmt.Println(spidermanPointer)

}