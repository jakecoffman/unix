unix
====

Implementation of Unix utilities in Go (Golang) for use as standalone binaries as well as libraries.

Goals of this repo
-----

- Have cross-platform implementations of Unix utilities
- Be able to use Unix pipelines "embedded" in other projects

Unix pipelines make it easy to do complex tasks with a one-liner. We could see similar benefits
if we could invoke them from within our program without having to run external commands. Especially
if you were running on a non-Unix server and the binaries aren't available.
