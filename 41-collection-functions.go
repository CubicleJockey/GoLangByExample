/*
We often need our programs to perform operations on collections of data,
like selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function.

In some languages it’s idiomatic to use generic data structures and algorithms.
Go does not support generics; in Go it’s common to provide collection functions if and when they are specifically
needed for your program and data types.

Here are some example collection functions for slices of strings.
You can use these examples to build your own functions.
Note that in some cases it may be clearest to just inline the collection-manipulating code directly,
instead of creating and calling a helper function.
*/
package main

import (
    "fmt"
    "strings"
)

//Index returns the first index of the target string, or -1 if no match is found.
func Index(values []string, target string) int {
    for index, value := range values {
        if value == target {
            return index
        }
    }
    return -1
}

//Include returns true if the target string is in the slice
func Include(values []string, target string) bool {
    return Index(values, target) >= 0
}

//Any returns true if one of the strings in the slice satisfies the predicate.
func Any(values []string, predicate func(string) bool) bool {
    for _, value := range values {
        if predicate(value) {
            return true
        }
    }
    return false
}

//All returns true if all of the strings in the slice satisfy the predicate.
func All(values []string, predicate func(string) bool) bool {
    for _, value := range values {
        if !predicate(value) {
            return false
        }
    }
    return true
}

//Filter returns a new slice containing all strings in the slice that satisfy the predicate.
func Filter(values []string, predicate func(string) bool) []string {
    predicateValues := make([]string, 0)
    for _, value := range values {
        if predicate(value) {
            predicateValues = append(predicateValues, value)
        }
    }
    return predicateValues
}

//Map returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(values []string, function func(string) string) []string {
    vsm := make([]string, len(values))
    for i, v := range values {
        vsm[i] = function(v)
    }
    return vsm
}

func main() {

    strs := []string{"peach", "apple", "pear", "plum"}
    fmt.Println(Index(strs, "pear"))
    fmt.Println(Include(strs, "grape"))
    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))
    fmt.Println(Map(strs, strings.ToUpper))

}