package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jakecoffman/unix"
)

func main() {
	in := make(chan string)
	out := unix.Grep(in, os.Args[1])

	go func() {
		defer func() { close(in) }()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			in <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		line, ok := <-out
		if !ok {
			break
		}
		fmt.Println(line)
	}
}
