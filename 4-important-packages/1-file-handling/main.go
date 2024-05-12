package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const filename = "file.txt"

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Create file
	////////////////////////////////////////////////////////////////////////////////////////////////
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Write to file
	////////////////////////////////////////////////////////////////////////////////////////////////
	message := "1234567890abcdefghijklmnopqrstuvwxyz\n" // file content
	// bytesWritten, err := file.WriteString(message)
	byteSlice := []byte(message)               // convert string to byte slice
	bytesWritten, err := file.Write(byteSlice) // same as WriteString, but you can write other types of data
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("File created successfully! Size: %d bytes\n", bytesWritten)

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Read entire file
	////////////////////////////////////////////////////////////////////////////////////////////////
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File Content:", string(fileContent))

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Read the file in chunks (obs: when the file is too big, it's better to read it in chunks)
	////////////////////////////////////////////////////////////////////////////////////////////////
	file2, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file for reading:", err)
		return
	}
	defer file2.Close()

	reader := bufio.NewReader(file2)
	// fmt.Println(reader.ReadString('\n'))
	buffer := make([]byte, 5) // create a buffer to read the file in chunks, in this case, 5 bytes at a time
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n])) // convert the buffer to a string and print it
	}

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Remove file
	////////////////////////////////////////////////////////////////////////////////////////////////
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
	fmt.Println("File removed successfully.")
}

func ReadFileInChunks() {

}

// go run 4-important-packages/1-file-handling/main.go
