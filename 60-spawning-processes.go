/* Sometimes our Go programs need to spawn other, non-Go processes. */

package main

import (
    "os/exec"
    "fmt"
    "io/ioutil"
)

func main() {

    dateCommand := exec.Command("date")

    dateOutput, err := dateCommand.Output()
    if err != nil {
        panic(err)
    }

    fmt.Println("> date")
    fmt.Println(string(dateOutput))

    grepCommand := exec.Command("grep", "hello")

    grepIn, _ := grepCommand.StdinPipe()
    grepOut, _ := grepCommand.StdoutPipe()

    grepCommand.Start()

    grepIn.Write([]byte("hello greap\ngoodbye grep"))
    grepIn.Close()

    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCommand.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    listCommand := exec.Command("bash", "-c", "ls -a -l -h")

    listOutput, err := listCommand.Output()
    if err != nil {
        panic(err)
    }

    fmt.Println("> ls -a -l -h")
    fmt.Println(string(listOutput))
}
