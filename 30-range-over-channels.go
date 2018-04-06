/*
In a previous example we saw how for and range provide iteration over basic data structures.
We can also use this syntax to iterate over values received from a channel.
*/

package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "Item 1"
    queue <- "Item 2"

    close(queue)

    for item := range queue {
        fmt.Println(item)
    }
}