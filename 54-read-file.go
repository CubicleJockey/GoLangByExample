package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//read's start at root. i.e windows C:\
	const PATH = "/tmp/dat.txt"

	//Reads entire contents
	content, err := ioutil.ReadFile(PATH)
	check(err)
	fmt.Println("File content:", string(content))
	fmt.Println()

	//os provides more control over file read
	file, err := os.Open(PATH)
	check(err)

	first5Bytes := make([]byte, 5)
	firstTotalRead, err := file.Read(first5Bytes)
	check(err)
	fmt.Printf("%d bytes: %s\n", firstTotalRead, string(first5Bytes))

	//jump(seek) to a known location in a file
	seekPosition, err := file.Seek(6, 0)
	check(err)

	//read from location seek was placed
	seekBytes := make([]byte, 2)
	totalSeekRead, err := file.Read(seekBytes)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", totalSeekRead, seekPosition, string(seekBytes))

	/*
	The io package provides some functions that may be helpful for file reading.
	For example, reads like the ones above can be more robustly implemented with ReadAtLeast.
	*/
	seekPosition2, err := file.Seek(6, 0)
	check(err)
	robustSeekBytes := make([]byte, 2)
	totalRobustSeekRead, err := io.ReadAtLeast(file, robustSeekBytes, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", totalRobustSeekRead, seekPosition2, string(robustSeekBytes))

	//Reset seek position back to beginning of file
	_, err = file.Seek(0, 0)
	check(err)

	/*
	The bufio package implements a buffered reader that may be useful both for its efficiency with many small reads
	and because of the additional reading methods it provides.
	*/

	bufferedReader := bufio.NewReader(file)
	peekBytes, err := bufferedReader.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(peekBytes))

	file.Close()
}

/*
Sample text to copy into a file on \\tmp\dat.text
-------------------------------------------------


Lorem Ipsum is simply dummy text of the printing and typesetting industry.
Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.
It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.
It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

*/
