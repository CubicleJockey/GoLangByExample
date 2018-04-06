/*
Basic sends and receives on channels are blocking.
However, we can use select with a default clause to implement non-blocking sends,
receives, and even non-blocking multi-way selects.
*/
package main

import "fmt"

func main(){

    messages := make(chan string)
    signals := make(chan bool)

    select {
        case message := <- messages:
            fmt.Println("Received message: ", message)
        default:
            fmt.Println("No messages recevied.")
    }

    message := "Hello"

    select {
        case messages <- message:
            fmt.Println("Sent Message: ", message)
        default:
            fmt.Println("No message sent.")
    }

    select {
        case message := <- messages:
            fmt.Println("Received message: ", message)
        case signal := <- signals:
            fmt.Println("Received signal: ", signal)
        default:
            fmt.Println("No acitivity")
    }

}
