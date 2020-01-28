package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

/*
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

//WordCount .
func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	finalData := map[string]int{}
	for _, e := range arr {
		finalData[e]++
	}

	return finalData
}

func main() {
	wc.Test(WordCount)
}
