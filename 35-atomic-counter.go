/*
The primary mechanism for managing state in Go is communication over channels.
We saw this for example with worker pools.
There are a few other options for managing state though.
Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
*/

package main

import (
    "sync/atomic"
    "time"
    "fmt"
)

func main() {

    var counter uint64

    for i := 0; i < 50; i++ {

        //generate a new anonymous function
        go func(){
           for {
               atomic.AddUint64(&counter, 1)
               time.Sleep(time.Millisecond)
           }
        }()
    }

    time.Sleep(time.Second)

    finalCount := atomic.LoadUint64(&counter)
    fmt.Println("Counter:", finalCount)
}