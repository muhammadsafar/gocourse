package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main80() {

	reader := bufio.NewReader(strings.NewReader("Hello , learn bufio package tesst!\n Words after enter"))

	//reading byte slice

	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		println("Error reading data:", err.Error())
		return
	}
	fmt.Printf("Read %d bytes: %s", n, data[:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		println("Error reading line:", err.Error())
		return
	}

	fmt.Println("\nRead line:", line)

	//* intinya buffer sebagai penampung data sementara sebelum di proses
	// pertama ambil slice 20 byte
	//sisa nya ambil sampai batas ketemu enter / new line

	//bufio.writer

	writer := bufio.NewWriter(os.Stdout)

	data2 := []byte("\nHello , write using bufio writer!\n")

	val, err := writer.Write(data2)

	if err != nil {
		fmt.Println("Error writing data:", err.Error())
		return
	}

	fmt.Printf("Wrote %d bytes", val)

	//flush the buffer to ensure all data is written
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing data:", err.Error())
		return
	}

	//writing string

	str := "Writing string using bufio writer!\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err.Error())
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	err = writer.Flush()

	if err != nil {
		fmt.Println("Error flushing data:", err.Error())
		return
	}

}
