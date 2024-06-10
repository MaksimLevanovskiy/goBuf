package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type bytesClosingBuffer struct {
	*bytes.Buffer
	io.Closer
}

func newBytesClosingBuffer() *bytesClosingBuffer {
	return &bytesClosingBuffer{
		Buffer: new(bytes.Buffer),
	}
}

func (b *bytesClosingBuffer) Close() error {
	fmt.Println("Closing")
	return nil
}

func writeTo(w io.WriteCloser, msg []byte) error {
	defer w.Close()

	_, err := w.Write(msg)
	return err
}

func main() {
	buf := newBytesClosingBuffer()

	err := writeTo(buf, []byte("New bytes"))
	if err != nil {
		log.Fatal("Errrrrr")
		//close
	}

	fmt.Println(buf.String())
}
