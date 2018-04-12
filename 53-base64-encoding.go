package main

import (
    "fmt"
    b64 "encoding/base64"
)

func main() {
    data := "abc123!?$*&()'-=@~"

    fmt.Println("Data to encode:", data)

    encodedString := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("Data encoded;", encodedString)

    //Decoding can return an error, ignored in this example
    decodedByteArray, _ := b64.StdEncoding.DecodeString(encodedString)
    fmt.Println("Decoded Byte Array:", decodedByteArray)
    fmt.Println("Decoded as string:", string(decodedByteArray))
    fmt.Println()

    encodedFromByteArray := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println("Encoded from byte array:", encodedFromByteArray)

    decodedFromByteArray, _ := b64.URLEncoding.DecodeString(encodedFromByteArray)
    fmt.Println("Decoded Byte Array from Byte Array:", decodedByteArray)
    fmt.Println("Decoded Byte Array from Byte Array as string:", string(decodedFromByteArray))
}