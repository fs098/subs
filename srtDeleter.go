package main

import (
	"fmt"
	"os"
	"strings"
)

// srtDeleter deletes all .srt files from a directory,
// outputs number of files deleted this way
func srtDeleter(dir string) {
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

		err := os.Remove(getPath(dir, file.Name()))
		if err != nil {
			fmt.Println(err)
		}
		delFileCount++
	}

	if delFileCount == 0 {
		fmt.Println("No files deleted.")
	} else {
		fmt.Println("Deleted", delFileCount, "files!")
	}
}

// srtDeleterR calls DeleterR function and outputs the number of files it deleted
func srtDeleterR(dir string) {
	delFileCount := DeleterR(dir)
	if delFileCount == 0 {
		fmt.Println("No files deleted.")
	} else {
		fmt.Println("Deleted", delFileCount, "files!")
	}
}

// DeleterR deletes all .srt files from a directory and it's subdirectories,
// returns number of files deleted this way
func DeleterR(dir string) int {
	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	delFileCount := 0
	for _, file := range dirContent {
		if file.IsDir() {
			delFileCount += DeleterR(getPath(dir, file.Name()))
			continue
		}
		isSubtitle := strings.HasSuffix(file.Name(), ".srt")
		if !isSubtitle {
			continue
		}

		err := os.Remove(getPath(dir, file.Name()))
		if err != nil {
			fmt.Println(err)
		}
		delFileCount++

	}
	return delFileCount
}
