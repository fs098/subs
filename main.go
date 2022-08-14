package main

import (
	"os"
)

func main() {
	flagsOrDir := os.Args[1:]
	srtDeleter(flagsOrDir[0])
}
