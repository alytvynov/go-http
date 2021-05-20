package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println(resp.StatusCode)

	io.Copy(os.Stdout, resp.Body)

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
