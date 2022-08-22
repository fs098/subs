package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	deleteFlag := flag.Bool("d", false, "Delete .srt files in directory")
	deleteRFlag := flag.Bool("dd", false, "Delete .srt files in directory and subdirectories")
	// selectIndex allows to manually select the index of the wanted .srt file. Starts at 1 instead of 0
	selectIndex := flag.Bool("i", false, "Fetches subtitles at given index")

	flag.Parse()
	myArgs := flag.Args()

	// myIndex is the index given to strCopy function, defaults to 0 but can be changed
	// with selectIndex flag. Ex: subs -i 2
	myIndex := 0

	var directories []string
	if !(*selectIndex) {
		directories = myArgs
	} else {
		intGiven, err := strconv.Atoi(myArgs[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		myIndex = intGiven - 1
		directories = myArgs[1:]
	}

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
		applyFlags(dir, myIndex, *deleteFlag, *deleteRFlag)
	}
}

func applyFlags(dir string, index int, delete bool, deleteR bool) {
	defaultB := !delete && !deleteR
	if defaultB {
		srtCopy(dir, index)
	} else if delete {
		srtDeleter(dir)
	} else if deleteR {
		srtDeleterR(dir)
	}
}

// TO DO:
// 1) remove overused os.ReadDir function
// 2) change copy function not to rely on bash
