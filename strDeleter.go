package main

import (
	"fmt"
	"os"
)

// srtDeleter deletes all .srt files from a directory and it's subdirectories,
// returns number of files deleted this way
func srtDeleter(dir string) int {
	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	delFileCount := 0
	for _, file := range dirContent {
		if file.IsDir() {
			delFileCount += srtDeleter(dir + "/" + file.Name())
		} else {
			start := len(file.Name()) - 4

			// fmt.Println(file.Name()[start:])
			if file.Name()[start:] == ".srt" {
				// fmt.Println(file.Name())
				err := os.Remove(dir + "/" + file.Name())
				if err != nil {
					fmt.Println(err)
				}
				delFileCount++
			}
		}
	}
	return delFileCount
}

func test() {
	flagsOrDir := os.Args[1:]
	filesDeleted := srtDeleter(flagsOrDir[0])
	fmt.Printf("%d files deleted\n", filesDeleted)
}
