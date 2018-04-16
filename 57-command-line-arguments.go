package main

import (
    "os"
    "fmt"
)

/*
1. To run this program in GoLand open the Terminal window.
2. go build 57-command-line-arguments
3. ./57-command-line-arguments a b c d
*/

func main() {

    argsWithProg := os.Args
    fmt.Println("Arguments with Program:", argsWithProg)

    argWithoutProg := os.Args[1:]
    fmt.Println("Arguments without Program:", argWithoutProg)

    specificArgument := os.Args[3]
    fmt.Println("Specific Argument:", specificArgument)
}