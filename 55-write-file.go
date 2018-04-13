package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

func panicOnError(e error){
	if e != nil {
		panic(e)
	}
}

func main() {

	const contents = "hello\ngo\n"
	const FILE1 = "/tmp/data-write.txt"

	fmt.Printf("Writing <\n%s> to file: [%s]", contents, FILE1)
	data := []byte(contents)
	err := ioutil.WriteFile(FILE1, data, 0644)
	panicOnError(err)

	fmt.Println()

	//create empty file
	file, fileError := os.Create("/tmp/data-write2.txt")
	panicOnError(fileError)

	defer file.Close()

	data2 := []byte{115, 111, 109, 101, 10}
	bytesWrittenCount, err := file.Write(data2)
	panicOnError(err)
	fmt.Printf("Wrote '%d' bytes: [%s]", bytesWrittenCount, string(data2))

	const WRITES = "Writes\n"
	bytesWrittenCount2, err := file.WriteString(WRITES)
	panicOnError(err)
	fmt.Printf("WroteString '%d' bytes: [%s]", bytesWrittenCount2, WRITES)

	//Writes to stable storage
	file.Sync()

	//buffered writer
	const BUFFERED = "Buffered\n"
	writer := bufio.NewWriter(file)
	bufferedBytesWritten, err := writer.WriteString(BUFFERED)
	panicOnError(err)
	fmt.Printf("Wrote buffered '%d' bytes: [%s]", bufferedBytesWritten, BUFFERED)

	//Apply all buffers to underlying writer
	writer.Flush()




}
