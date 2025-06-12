//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := os.Args[1:]
	signals := make(chan string)

	for _, u := range urls {
		go func(u string) {
			start := time.Now()
			resp, err := http.Get(u)
			if err != nil {
				fmt.Printf("%v\n", err)
				signals <- fmt.Sprintf("%v", err)
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%v\n", err)
				signals <- fmt.Sprintf("%v", err)
				return
			}

			elapsed := time.Since(start)
			signals <- fmt.Sprintf("%v\t%v\t%v", elapsed, len(body), u)
		}(u)
	}

	for range urls {
		msg := <-signals
		fmt.Println(msg)
	}
}
