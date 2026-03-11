package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesFromFile(file io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer file.Close()
		defer close(out)

		var line string
		for {
			data := make([]byte, 8)
			n, err := file.Read(data)
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
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	lines := getLinesFromFile(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}

}
