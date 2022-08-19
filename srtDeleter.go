package main

import (
	"fmt"
	"os"
	"strings"
)

// srtDeleter deletes all .srt files from a directory,
// returns number of files deleted this way
func srtDeleter(dir string) int {
	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	delFileCount := 0
	for _, file := range dirContent {
		isSubtitle := strings.HasSuffix(file.Name(), ".srt")
		if !isSubtitle {
			continue
		}

		err := os.Remove(dir + "/" + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		delFileCount++
	}
	return delFileCount
}

// srtDeleterR deletes all .srt files from a directory and it's subdirectories,
// returns number of files deleted this way
func srtDeleterR(dir string) int {
	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	delFileCount := 0
	for _, file := range dirContent {
		if file.IsDir() {
			delFileCount += srtDeleterR(dir + "/" + file.Name())
			continue
		}
		isSubtitle := strings.HasSuffix(file.Name(), ".srt")
		if !isSubtitle {
			continue
		}

		err := os.Remove(dir + "/" + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		delFileCount++

	}
	return delFileCount
}

func test() {
	flagsOrDir := os.Args[1:]
	filesDeleted := srtDeleter(flagsOrDir[0])
	remainigFiles := srtDeleterR(flagsOrDir[0])
	fmt.Printf("%d files deleted + %d files in folders\n", filesDeleted, remainigFiles)
}
