package unix

import "io/ioutil"

// TODO: Support stdin optionally
func Ls() <-chan string {
	out := make(chan string)

	go func() {
		defer func() {
			close(out)
		}()

		files, err := ioutil.ReadDir(".")
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			out <- file.Name()
		}
	}()

	return out
}
