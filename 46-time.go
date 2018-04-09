/*
Go offers extensive support for times and durations; here are some examples.
*/

package main

import (
    "fmt"
    "time"
)

func main() {

    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p("Year:", then.Year())
    p("Month:", then.Month())
    p("Day:", then.Day())
    p("Hour:", then.Hour())
    p("Minute:", then.Minute())
    p("Second:", then.Second())
    p("Nanosecond:", then.Nanosecond())
    p("Location:", then.Location())
    p("Weekday:", then.Weekday())

    p("Before:", then.Before(now))
    p("After:",then.After(now))
    p("Equal", then.Equal(now))

    diff := now.Sub(then)
    p("Difference between ", then, "and", now, "is", diff)
    p("Difference in hours:",diff.Hours())
    p("Difference in minutes:", diff.Minutes())
    p("Difference in seconds:", diff.Seconds())
    p("Difference in nanoseconds:", diff.Nanoseconds())

    p("Adding difference to then:", then.Add(diff))
    p("Subtracting difference to then:", then.Add(-diff))
}