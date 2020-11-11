package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	maps := make(map[string]int)
	for _, word := range words {
		maps[word]++
	}
	return maps
}

func main() {
	wc.Test(WordCount)
}
