/*

A line filter is a common type of program that reads input on stdin,
processes it, and then prints some derived result to stdout.

grep and sed are common line filters.

*/

package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {

	/*
	Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the
	scanner to the next token; which is the next line in the default scanner.
	*/
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := strings.ToUpper(scanner.Text())
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
