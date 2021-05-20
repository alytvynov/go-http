package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	/*
		bs := make([]byte, 999999)
		len, readErr := resp.Body.Read(bs)

		fmt.Println(len)
		fmt.Println(string(bs))

		if readErr != nil && readErr != io.EOF {
			fmt.Println("Read Error", readErr)
			os.Exit(1)
		}
	*/
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))

	return len(bs), nil
}
