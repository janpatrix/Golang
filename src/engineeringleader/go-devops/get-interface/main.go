package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	contents string
	pos      int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 < len(m.contents) {
		n := copy(p, []byte(m.contents[m.pos:m.pos+1]))
		m.pos++
		return n, nil
	}
	return n, io.EOF
}

func main() {

	mySlowReader := &MySlowReader{
		contents: "hello world!",
		pos:      0,
	}

	out, err := io.ReadAll(mySlowReader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %s\n", out)
}
