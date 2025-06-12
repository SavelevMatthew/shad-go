//go:build !solution

package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]
	counters := make(map[string]int64)
	for _, fileName := range args {
		data, err := os.ReadFile(fileName)
		check(err)
		stringData := string(data)
		fields := strings.Split(stringData, "\n")

		for _, f := range fields {
			counters[f]++
		}
	}

	for key, value := range counters {
		if value > 1 {
			fmt.Printf("%v\t%v\n", value, key)
		}
	}
}
