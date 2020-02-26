package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Not enough args")
	}

	if !strings.HasPrefix(args[0], "-depth") {
		panic("Not set depth")
	}

	data := strings.Split(args[0], "=")
	DEPTH, err := strconv.Atoi(data[1])

	if err != nil {
		panic("Error parsing")
	}

	worklist := make(chan []string)
	go func() { worklist <- args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for i := 0; i <= DEPTH; i++ {
		listOfLinks := <-worklist
		for _, link := range listOfLinks {
			if !seen[link] {
				seen[link] = true

				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}

		fmt.Println("LEVEL: " + strconv.Itoa(i))

	}
}

//!-
