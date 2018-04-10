/*
    Parsing numbers from strings is a basic but common task in many programs; here’s how to do it in Go.
    The built-in package strconv provides the number parsing.
*/
package main

import (
    "strconv"
    "fmt"
)

func main() {

    float64, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println("Parsed Float64:", float64)

    integer64, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println("Parsed Integer64:", integer64)

    integer64FromHex, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println("Parsed Integer64 from Hex:", integer64FromHex)

    unsignedInteger64, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println("Parsed Unsigned Integer64", unsignedInteger64)

    base10Integer, _ := strconv.Atoi("135")
    fmt.Println("Parsed base10 Integer:", base10Integer)

    _, error := strconv.Atoi("not an integer")
    fmt.Println("Error:", error)
}
