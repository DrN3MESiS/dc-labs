package main

import (
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("pingpong.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log.SetFlags(0)
	log.SetOutput(f)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.Printf("Start Time: " + time.Now().String())

	chanPing := make(chan int)
	chanPong := make(chan int)

	for true {
		<-chanPing
		go func(c chan int) {
			log.Printf("ping -> " + time.Now().String())
			c <- 1
		}(chanPong)

		<-chanPong
		go func(c chan int) {
			log.Printf("pong -> " + time.Now().String())
			c <- 1
		}(chanPing)
	}
	log.Printf("Finished")
}
