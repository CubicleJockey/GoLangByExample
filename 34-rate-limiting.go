/*

Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
Go elegantly supports rate limiting with goroutines, channels, and tickers.

*/

package main

import (
    "time"
    "fmt"
)

func main(){

    const MAX = 5;

    requests := make(chan int, MAX)
    for i := 1; i <= MAX; i++ {
        requests <- i
    }

    delay := 200 * time.Millisecond
    limiter := time.Tick(delay)

    //Retrieve a request every 200 milliseconds
    for request := range requests {
        <- limiter
        fmt.Println("Request", request, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3)

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for tick := range time.Tick(delay) {
            burstyLimiter <- tick
        }
    }()

    burstyRequests := make(chan int, MAX)
    for i := 1; i <= MAX; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for request := range burstyRequests {
        <- burstyLimiter
        fmt.Println("Bursty Request:", request, time.Now())
    }
    fmt.Scanln()
    fmt.Println("done.")
}