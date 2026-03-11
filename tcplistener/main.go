package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesFromChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		var line string
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			if s := string(data[:n]); strings.Contains(s, "\n") {
				before, after, _ := strings.Cut(s, "\n")
				line += before
				out <- line
				line = after
			} else {
				line += s
			}
		}
	}()

	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("err", "err", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("err", "err", err)
		}
		for line := range getLinesFromChannel(conn) {
			fmt.Printf("read: %s\n", line)
		}
	}

}
