package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading data from reader", err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing data to writer", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error close", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer //stact
	buf.WriteString("Hello String...")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) //heap
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading buffer", err)
	}

	fmt.Println(buf.String())

}

func pipeExample() {

	pr, pw := io.Pipe()

	go func() {
		pw.Write([]byte(`hello pipe`))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)

	fmt.Println(buf.String())
}

func writeToFile(filePath string, data string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error opening/creating file", err)
	}

	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("Error opening/creating file", err)
	// }

}

func main99() {

	fmt.Println("======READ FROM READER=========")
	readFromReader(strings.NewReader("Hello reader!"))

	fmt.Println("======WRITE TO WRITER=========")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello writer")
	fmt.Println(writer.String())

	fmt.Println("======BUFFER EXAMPLE=========")
	bufferExample()

	fmt.Println("======Multi Reader EXAMPLE=========")
	multiReaderExample()

	fmt.Println("======Pipe EXAMPLE=========")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "Hello file")

	resource := &MyResource{name: "Test Resource"}
	closeResource(resource)

}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Println("Closing resource", m.name)
	return nil
}
