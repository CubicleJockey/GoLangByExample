/*
A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch.
Here’s how to do it in Go.
*/
package main

import (
    "time"
    "fmt"
)

func main() {

    now := time.Now()
    fmt.Println("Now:", now)

    seconds := now.Unix()
    fmt.Println("Seconds:", seconds)

    nanoseconds := now.UnixNano()
    fmt.Println("Nano-Seconds:", nanoseconds)

    //There is no milliseconds so you'll manually have to divide
    milliseconds := nanoseconds / 10000
    fmt.Println("Milliseconds:", milliseconds)


    fmt.Println(time.Unix(seconds, 0))
    fmt.Println(time.Unix(0, nanoseconds))
}
