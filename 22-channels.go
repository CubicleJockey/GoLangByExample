/*

Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.

Create a new channel with make(chan val-type). Channels are typed by the values they convey.

Send a value into a channel using the channel <- syntax.

The <-channel syntax receives a value from the channel.

By default sends and receives block until both the sender and receiver are ready.
This property allowed us to wait at the end of our program for the
"ping" message without having to use any other synchronization.

*/

package main

import "fmt"

func main() {

    messages := make(chan string)

    //Here we send "ping" to the messages channel we made above, from a new goroutine.
    go func() {messages <- "ping"}()


    //Here we’ll receive the "ping" message we sent above and print it out.
    channelResponse := <- messages
    fmt.Println(channelResponse)

}
