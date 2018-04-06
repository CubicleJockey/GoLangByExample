/*

Closing a channel indicates that no more values will be sent on it.
This can be useful to communicate completion to the channel’s receivers.

*/

package main

import "fmt"

func main(){

    jobs := make(chan int, 5)
    done := make(chan bool)

    // anonymous job that loops until all jobs are done.
    go func(message string){

        fmt.Println(message)

        for {
            job, more := <- jobs
            if more {
                fmt.Println("Recieved job", job)
            } else {
                fmt.Println("Received all jobs.")
                done <- true
                fmt.Println("Done.")

                return
            }
        }
    }("Getting jobs...")

    for index := 1; index <= 5; index++ {
        jobs <- index
        fmt.Println("Sent job", index)
    }

    close(jobs)
    fmt.Println("Sent all jobs")

}
