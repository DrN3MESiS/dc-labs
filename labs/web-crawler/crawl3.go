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

	list := args[1:]
	curDEPTH := 0
	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for true {
		if curDEPTH > DEPTH {
			break
		}
		fmt.Println("Level:" + strconv.Itoa(curDEPTH))

		DEPTHlist := [][]string{}

		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				fmt.Println("\t" + link)
				go func(link string) {
					worklist <- crawl(link)
				}(link)
				DEPTHlist = append(DEPTHlist, <-worklist)
			}
		}

		list = []string{}

		for _, li := range DEPTHlist {
			for _, link := range li {
				list = append(list, link)
			}
		}

		curDEPTH++
	}

}

//!-
