/*

Timers are for when you want to do something once in the future.
Tickers are for when you want to do something repeatedly at regular intervals.

*/
package main

import (
    "time"
    "fmt"
)

func main(){

    ticker := time.NewTimer(50 * time.Millisecond)

    go func(){
        //get ticks as they come in from the ticker channel.
        for tick := range ticker.C {
            fmt.Println("Tick at: ", tick)
        }
    }()

    time.Sleep(10 * time.Second)

    ticker.Stop()
    fmt.Println("Ticker set to stop.")
}