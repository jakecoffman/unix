package main

import (
	"fmt"

	"github.com/jakecoffman/unix"
)

func main() {
	out := unix.Ls()
	for {
		line, ok := <-out
		if !ok {
			break
		}
		fmt.Println(line)
	}
}
