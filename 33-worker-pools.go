package main

import (
    "fmt"
    "time"
)

type Result struct {
    Id int
    Job int
    Result int
}

func worker(id int, jobs <- chan int, results chan <- Result) {
    for job := range jobs {
        fmt.Println("Worker", id, "started job", job)
        time.Sleep(2 * time.Second)
        fmt.Println("Worker", id, "finished job", job)
        results <- Result{id, job, job * 2 }
    }
}

func main() {

    const TOTAL = 100
    const JOB_COUNT = 7

    jobs := make(chan int, TOTAL)
    results := make(chan Result, TOTAL)

    for work := 1; work <= 5; work++ {
        go worker(work, jobs, results)
    }

    for job := 1; job <= JOB_COUNT; job++ {
        jobs <- job
    }

    close(jobs)

    for i := 1; i <= JOB_COUNT; i++ {
        result := <- results
        fmt.Println("Worker Id:", result.Id, "Job:", result.Job, "Result:", result.Result)
    }

    close(results)

}
