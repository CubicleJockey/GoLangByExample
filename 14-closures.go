/*
Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function inline without having to name it.
*/

package main

import "fmt"

/*
This function integerSequence returns another function, which we define anonymously in the body of intSeq.
The returned function closes over the variable i to form a closure.
*/
func integerSequence() func() int {
    var i int
    return func() int {
        i++
        return i
    }
}

func main(){

    const originalName = "nextInt: "
    const freshName = "freshStart: "

    nextInt := integerSequence()

    fmt.Println(originalName, nextInt())
    fmt.Println(originalName, nextInt())
    fmt.Println(originalName, nextInt())


    freshStart := integerSequence()
    fmt.Println(freshName, freshStart())

    fmt.Println(originalName, nextInt())

    fmt.Println(freshName, freshStart())


}