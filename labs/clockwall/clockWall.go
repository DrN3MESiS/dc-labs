package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	c := make(chan int)
	for _, url := range os.Args[1:] {
		data := strings.Split(url, "=")

		go getData(c, data)
	}
	data := <-c
	fmt.Println(data)
}

func getData(c chan int, data []string) {
	fmt.Println(data)
	conn, err := net.Dial("tcp", data[1])
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	io.Copy(os.Stdout, conn)
	c <- 1
}
