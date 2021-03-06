/*

We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.

*/

package main

import (
    "fmt"
    "time"
)

/*
This is the function we’ll run in a goroutine.
The done channel will be used to notify another goroutine that this function’s work is done.
*/
func worker(done chan bool) {
    fmt.Println("Starting work...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <- done //blocked until value is set
}