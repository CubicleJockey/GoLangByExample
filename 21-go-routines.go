/* A goroutine is a lightweight thread of execution. */
package main

import "fmt"

func DoThingy(from string) {
    for i := 0; i < len(from); i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    DoThingy("synchronous")

    go DoThingy("asynchronous")

    //anonymous async to intertwine with go DoThingy
    go func(message string){
        fmt.Println(message)
   }("anonymous")

    fmt.Scanln() //<--- get async outputs
    fmt.Println("done")
}