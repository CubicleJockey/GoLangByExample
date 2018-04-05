/*
By default channels are unbuffered.
Meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/

package main

import "fmt"

func main() {

    //Here we make a channel of strings buffering up to 2 values.
    messages := make(chan string, 3)

    //Non-Blocking channel value sets
    messages <- "buffered" //1
    messages <- "channel"  //2
    messages <- "after doing something" //3

    fmt.Println(<-messages)
    fmt.Println(<-messages)

    fmt.Println("Doing something else before getting 3rd buffered channel value")

    fmt.Println(<-messages)
}
