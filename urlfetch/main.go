//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		resp, err := http.Get(url)
		check(err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		check(err)
		s := string(body)
		fmt.Println(s)
	}
}
