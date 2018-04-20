package main

import (
    "os/exec"
    "os"
    "syscall"
)

func main() {

    const lsCommand = "ls"

    binary, lookErr := exec.LookPath(lsCommand)
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string { lsCommand, "-a", "-l", "-h" }

    //get environment so exec can use appropriate environment variables
    environment := os.Environ()

    //output of syscall.Exec will go to stdout
    executeError := syscall.Exec(binary, args, environment)
    if executeError != nil {
        panic(executeError)
    }
}
