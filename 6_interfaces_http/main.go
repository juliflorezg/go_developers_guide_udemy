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
		fmt.Println("Error on get petition:", err)

		os.Exit(1)

	}

	lw := logWriter{}

	// byte slice
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	num, err := fmt.Println(string(bs))

	fmt.Println("Just wrote this many bytes:", num)

	return num, err
}
