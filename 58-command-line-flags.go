package main

import (
    "flag"
    "fmt"
)

/*
Command-line flags are a common way to specify options for command-line programs.
For example, in wc -l the -l is a command-line flag.

To run this:
1. Open GoLand terminal
2. go build 58-command-line-flags
3. ./58-command-line-flags -AFlag=SomeValue -IntFlag=333 -BoolFlag -sVariable=anotherFlag
*/

func main() {

    const AFLAG = "AFlag"
    commandFlagsPtr := flag.String(AFLAG, "foo", "a string")

    fmt.Println(AFLAG + " value:", *commandFlagsPtr)

    const INTFLAG = "IntFlag"
    intPtr := flag.Int(INTFLAG, 42, "an integer")

    fmt.Println(INTFLAG + " value:", *intPtr)

    const BOOLFLAG = "BoolFlag"
    boolPtr := flag.Bool(BOOLFLAG, false, "a boolean")

    fmt.Println(BOOLFLAG + " value:", *boolPtr)

    var sVariable string
    /*
    It’s also possible to declare an option that uses an existing var declared elsewhere in the program.
    Note that we need to pass in a pointer to the flag declaration function.
    */
    flag.StringVar(&sVariable, "sVariable", "bar", "A string variable")

    fmt.Println("sVariable value:", sVariable)

    //Once all flags are declared, call flag.Parse() to execute the command-line parsing
    flag.Parse()


    fmt.Println("Values being passed in via command-line as flags")
    fmt.Println()

    fmt.Println(AFLAG + ":", *commandFlagsPtr)
    fmt.Println(INTFLAG + ":", *intPtr)
    fmt.Println(BOOLFLAG + ":", *boolPtr)
    fmt.Println("sVariable:", sVariable)
    fmt.Println("Command-line arguments:", flag.Args())
}