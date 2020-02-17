package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		data := strings.Split(url, "=")
		done := make(chan string)
		go getData(done, data)
		out := <-done
		fmt.Println(out)
	}
}

func getData(done chan string, data []string) {
	conn, err := net.Dial("tcp", data[1])
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, conn)
	done <- fmt.Sprintf("%v", conn)
}
