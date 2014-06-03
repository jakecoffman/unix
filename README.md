unix
====

Implementation of Unix utilities in Go (Golang) for use as standalone binaries as well as libraries.

goals of this repo
-----

- Have cross-platform implementations of Unix utilities
- Be able to use Unix pipelines "embedded" in other projects

Unix pipelines make it easy to do complex tasks with a one-liner. We could see similar benefits
if we could invoke them from within our program without having to run external commands. Especially
if you were running on a non-Unix server and the binaries aren't available.

installing binaries
-------------------

There may be a better way to do this, but running these commands:

```
go get github.com/jakecoffman/unix/execs/ls
cd %GOPATH%\github.com\jakecoffman\unix\execs\ls
go install
```

Will install ls into $GOBIN on Windows.
