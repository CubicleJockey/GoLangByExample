/*
A panic typically means something went unexpectedly wrong.
Mostly we use it to fail fast on errors that shouldn’t occur during normal operation,
or that we aren’t prepared to handle gracefully.

Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use
error-indicating return values wherever possible.
*/

package main

import (
    "os"
)

func main() {

    panic("Oh noes n' stuff.")

    _, error := os.Create("/tmp/file")
    if error != nil {
        panic(error)
    }
}
