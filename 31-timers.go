/*

We often want to execute Go code at some point in the future, or repeatedly at some interval.
Go’s built-in timer and ticker features make both of these tasks easy. We’ll look first at timers and then at tickers.

*/

package main

import (
    "time"
    "fmt"
)

func main() {

    timer1 := time.NewTimer(2 * time.Second)

    //Block till timer channel is complete.
    <- timer1.C
    fmt.Println("Timer 1 expired.")

    // really long timer
    timer2 := time.NewTimer(67 * time.Second)

    go func(){
        //block until timer channel is complete.
        <- timer2.C
    }()

    //stop the timer despite run time
    stop2 := timer2.Stop()
    if stop2 {
       fmt.Println("Timer 2 stopped.")
    }

}