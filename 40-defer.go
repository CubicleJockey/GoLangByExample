/*
Defer is used to ensure that a function call is performed later in a program’s execution,
usually for purposes of cleanup.

defer is often used where e.g. ensure and finally would be used in other languages.
*/

package main

import (
    "os"
    "fmt"
)

func main() {

    /*
    Immediately after getting a file object with createFile, we defer the closing of that file with closeFile.
    This will be executed at the end of the enclosing function (main), after writeFile has finished.
    */
    file := createFile("/tmp/defer.txt")
    defer closeFile(file)
    writeFile(file)

}

func createFile(fullFilePath string) *os.File {
    fmt.Println("Creating file:", fullFilePath)
    file, err := os.Create(fullFilePath)
    if err != nil {
        panic(err)
    }
    return file
}

func writeFile(fullFilePath *os.File) {
    fmt.Println("Writing file:", fullFilePath)
    fmt.Fprintln(fullFilePath, "Inserting data.")
}

func closeFile(fullFilePath *os.File) {
    fmt.Println("Closing file:", fullFilePath)
    fullFilePath.Close()
}