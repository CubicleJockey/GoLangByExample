/*
Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
Implementing timeouts in Go is easy and elegant thanks to channels and select.
*/

package main

import (
    "time"
    "fmt"
)

func main() {

    channel1 := make(chan string, 1)

    go func(){
        time.Sleep(2 * time.Second)
        channel1 <- "Channel 1 Response"
    }()

    select {
        case response := <- channel1:
            fmt.Println(response)
        case <- time.After(1 * time.Second):
            fmt.Println("A timeout occured for channel 1.")
    }

    channel2 := make(chan string, 1)

    go func(){
        time.Sleep(2 * time.Second)
        channel2 <- "Channel 2 Response"
    }()

    select {
        case response := <- channel2:
            fmt.Println(response)
        case <- time.After(3 * time.Second):
            fmt.Println("A timeout occured for channel 2")
    }


}