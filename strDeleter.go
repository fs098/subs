package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// srtDeleter deletes all .srt files from a directory and it's subdirectories
func srtDeleter(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			srtDeleter(dir + "/" + fileInfo.Name())
		} else {
			start := len(fileInfo.Name()) - 4

			// fmt.Println(fileInfo.Name()[start:])
			if fileInfo.Name()[start:] == ".srt" {
				// fmt.Println(fileInfo.Name())
				err := os.Remove(dir + "/" + fileInfo.Name())
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
