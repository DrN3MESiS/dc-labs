package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	f, err := os.OpenFile("ping-pong.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log.SetFlags(0)
	log.SetOutput(f)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	t := time.Now()
	c := make(chan int)
	n := 0
	go inf(c, t, &n, false)
	d := <-c
	log.Printf("==================\nTime waited until complete: %s, %+v", time.Since(t).String(), d)
}

func inf(c chan int, t time.Time, n *int, dev bool) int {
	//Log each
	*n++
	log.Println(strconv.Itoa(*n) + " -> " + time.Since(t).String())
	//Make res channel
	res := make(chan int)

	if dev {
		go inf(res, t, n, dev)
	}

	if *n < 2000 {
		go inf(res, t, n, dev)
	}

	res <- 1

	return <-res

}
