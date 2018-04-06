/*

In the previous example we saw how to manage simple counter state using atomic operations.
For more complex state we can use a mutex to safely access data across multiple goroutines.

*/

package main

import (
    "sync"
    "math/rand"
    "sync/atomic"
    "time"
    "fmt"
)

func main() {

    const MAXKEY = 5
    const DELTA = 1

    var state = make(map[int]int)

    var mutex = &sync.Mutex{}

    var readOps uint64
    var writeOps uint64

    for read := 0; read < 200; read++ {
        go func(){
            total := 0
            for {
                key := rand.Intn(MAXKEY)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, DELTA)

                time.Sleep(time.Millisecond)
            }
        }()
    }

    for write := 0; write < 100; write++ {
        go func(){
            for {
                key := rand.Intn(MAXKEY)
                value := rand.Intn(10)
                mutex.Lock()
                state[key] += value
                mutex.Unlock()
                atomic.AddUint64(&writeOps, DELTA)

                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOperationsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("Read operations:", readOperationsFinal)

    writeOperationsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("Write Operations:", writeOperationsFinal)

    mutex.Lock()
    fmt.Println("State:", state)
    mutex.Unlock()

}