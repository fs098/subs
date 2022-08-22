package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// fetches .srt files by default
	deleteFlag := flag.Bool("d", false, "Delete .srt files in directory")
	deleteRFlag := flag.Bool("dd", false, "Delete .srt files in directory and subdirectories")
	// selectIndex allows to manually select the index of the wanted .srt file. Starts at 1 instead of 0
	// selectIndex := flag.Bool("i", false, "Fetches subtitles at given index")

	flag.Parse()
	directories := flag.Args()

	var myDirs []string
	if len(directories) == 0 {
		// if there are no args in directories then use the current path
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}

		myDirs = append(myDirs, currentDir)
	} else {
		for _, entry := range directories {
			// check if string given is a valid directory
			dir, err := os.Stat(entry)
			if err != nil {
				fmt.Println(err)
				return
			}
			if dir.IsDir() {
				myDirs = append(myDirs, entry)
			}
		}
	}

	for _, dir := range myDirs {
		applyFlags(dir, *deleteFlag, *deleteRFlag)
	}
}

func applyFlags(dir string, delete bool, deleteR bool) {
	defaultB := !delete && !deleteR
	if defaultB {
		srtCopy(dir, 0)
	} else if delete {
		srtDeleter(dir)
	} else if deleteR {
		srtDeleterR(dir)
	}
}

// TO DO:
// 1) remove overused os.ReadDir function
// 2) select SDH subtitle
// 3) change copy function not to rely on bash
// 4) teste srtDeleter, after "/" change
