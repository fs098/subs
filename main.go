package main

import (
	"fmt"
	"os"
)

func main() {
	flagsOrDir := os.Args[1:]
	filesDeleted := srtDeleter(flagsOrDir[0])
	fmt.Printf("%d files deleted\n", filesDeleted)
}
