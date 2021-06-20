package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// byteSlice := make([]byte, 99999) // create an empty byte slice with initial size of 99999.
	// resp.Body.Read(byteSlice) // if we pass in slice with size 0, reader gonna think that it's alreay full
	// fmt.Println(string(byteSlice))

	// using Go helper to simplify reading body logic
	// io.Copy(os.Stdout, resp.Body) // take in writer,reader interface's member

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
