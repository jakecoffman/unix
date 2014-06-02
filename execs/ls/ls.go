package main

import (
	"fmt"

	"github.com/jakecoffman/unix/ls"
)

func main() {
	out := ls.Ls()
	for {
		line, ok := <-out
		if !ok {
			break
		}
		fmt.Println(line)
	}
}
