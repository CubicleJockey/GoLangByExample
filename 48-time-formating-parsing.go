/*

Go supports time formatting and parsing via pattern-based layouts.

*/

package main

import (
    "fmt"
    "time"
)

func main() {

    stdout := fmt.Println
    format := fmt.Printf

    now := time.Now()
    stdout(now.Format(time.RFC3339))

    time1, error2 := time.Parse(time.RFC3339,"2012-11-01T22:08:41+00:00")
    stdout("Parsed Time:", time1)
    stdout("Error:", error2)

    stdout("Formatted time string:", now.Format("3:04PM"))
    stdout("Formatted date string:", now.Format("Mon Jan _2 15:04:05 2006"))
    stdout("Formatted date string:", now.Format("2006-01-02T15:04:05.999999-07:00"))

    timeForm := "3 04 PM"
    time2, error2 := time.Parse(timeForm, "8 41 PM")
    stdout("Parsed time:", time2)
    stdout("Error:", error2)

    format("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        now.Year(), now.Month(), now.Day(),
        now.Hour(), now.Minute(), now.Second())

    malformedDate := "Mon Jan _2 15:04:05 2006"
    _, malformedError := time.Parse(malformedDate, "8:41PM")
    stdout("Parse Error:", malformedError)
}