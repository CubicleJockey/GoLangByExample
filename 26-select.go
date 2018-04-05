/*
  Go’s select lets you wait on multiple channel operations.
  Combining goroutines and channels with select is a powerful feature of Go.
*/

package main

import (
    "time"
    "fmt"
)

func main() {

    channel1 := make(chan string)
    channel2 := make(chan string)

    //anonymous go-routine
    go func(){
       time.Sleep(1 * time.Second)
       channel1 <- "Channel One Response"
    }()

    //anonymous go-routine
    go func(){
        time.Sleep(4 * time.Second)
        channel2 <- "Channel Two Response"
    }()

    var totalChannels int
    fmt.Println("Begin wait for channels...")
    for {
        select {
            case message1 := <- channel1:
                fmt.Println("Received: ", message1)
                totalChannels++
            case message2 := <- channel2:
                fmt.Println("Received: ", message2)
                totalChannels++
        }
        if totalChannels == 2 {
            fmt.Println("All channels have returned.")
            break
        }
    }
}