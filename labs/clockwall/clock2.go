// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func timeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func handleConn(c net.Conn, d string) {
	defer c.Close()
	for {

		convTime, err := timeIn(time.Now(), d)
		_, err = io.WriteString(c, convTime.Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	param := os.Args[1]
	if param != "-port" {
		panic("bruh")
	}
	port := os.Args[2]

	url := fmt.Sprintf("localhost:%v", port)
	fmt.Println("A new server has been hosted on: http://" + url)
	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, os.Getenv("TZ")) // handle connections concurrently
	}
}
