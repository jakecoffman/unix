package grep

import "strings"

func Grep(in <-chan string, needle string) <-chan string {
	out := make(chan string)
	go func() {
		defer func() { close(out) }()

		for {
			line, ok := <-in
			if !ok {
				break
			}
			if strings.Contains(line, needle) {
				out <- line
			}
		}
	}()
	return out
}
