package main

import (
	"net/http"

	"github.com/jakecoffman/unix"
)

func main() {
	http.HandleFunc("/", Helper(func() <-chan string {
		pipe := unix.Ls()
		return unix.Grep(pipe, "test")
	}))
	http.ListenAndServe("localhost:8080", nil)
}

func Helper(f func() <-chan string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		out := f()
		for {
			line, ok := <-out
			if !ok {
				return
			}
			w.Write([]byte(line))
		}
	}
}
