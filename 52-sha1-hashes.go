package main

import (
    "crypto/sha1"
    "fmt"
)

func main() {

    strToHash := "I am going to be hashed"

    hash := sha1.New()

    hash.Write([]byte(strToHash))

    byteSlice := hash.Sum(nil)

    fmt.Println("Source:", strToHash)
    fmt.Printf("Source Hash: %x\n", byteSlice)

}
