/*
Go offers excellent support for string formatting in the printf tradition.
Here are some examples of common string formatting tasks.
*/

package main

import (
    "fmt"
    "os"
)

type Point struct {
    X int
    Y int
}

func main() {

    point := Point {1, 2}

    //place-holder
    fmt.Printf("%v\n", point)

    //place-holder that prints property names
    fmt.Printf("%+v\n", point)

    //go syntax representation of value
    fmt.Printf("%#v\n", point)

    //To print the type of a value, use %T.
    fmt.Printf("%T\n", point)

    //Formatting bool
    fmt.Printf("%t\n", true)

    //Formatting base 10 digits
    fmt.Printf("%d\n", 123)

    //Prints binary representation
    fmt.Printf("%b\n", 14)

    //Prints character representation of integer
    fmt.Printf("%c\n", 33)

    //prints hex representation
    fmt.Printf("%x\n", 456)

    //prints floats
    fmt.Printf("%f\n", 78.9)

    //slightly different version of scientific notation
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    //basic string print
    fmt.Printf("%s\n", "\"string\"")

    //To double-quote strings as in Go source, use %q.
    fmt.Printf("%q\n", "\"string\"")

    //Hex representation of string
    fmt.Printf("%x\n", "hex this")

    //Print representation of a pointer
    fmt.Printf("%p\n", &point)

    //control width
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    //control width and precision of floating points
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    //control width and precision and left justify floating  points
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    //control string widths
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    //control string widths with left justify
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    //output to a string instead of stdout
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    //format to stderr
    fmt.Fprintf(os.Stderr, "an %s\n", "error")


}