/*

Go’s math/rand package provides pseudo-random number generation.

*/
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

    print := fmt.Print
    printLine := fmt.Println

    print("Randomly generated - [", rand.Intn(100), ",", rand.Intn(100), "] <same each run>")
    printLine()

    printLine("Random float-64:", rand.Float64())

    // The default number generator is deterministic, so it’ll produce the same sequence of numbers each time by default.
    // To produce varying sequences, give it a seed that changes.
    // Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.
    newSequence := rand.NewSource(time.Now().UnixNano())
    varyingRandomGenerator := rand.New(newSequence)

    print("Randomly generated - [",
              varyingRandomGenerator.Intn(100),
              ",",
              varyingRandomGenerator.Intn(100),
              "] <different each run>")
    printLine()



}