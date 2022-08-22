package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// main defines the flags that can be used by the program,
// gets the index to be used by srtCopy function, the
// directory or directories to be used and checks if they are valid.
// finally, main calls applyFlags function to each directory.
func main() {
	deleteFlag := flag.Bool("d", false, "Delete .srt files in directory")
	deleteRFlag := flag.Bool("dd", false, "Delete .srt files in directory and subdirectories")
	// selectIndex allows to manually select the index of the wanted .srt file. Starts at 1 instead of 0
	selectIndex := flag.Bool("i", false, "Fetches subtitles at given index")

	flag.Parse()
	// myArgs should consist of directories or an int and directories
	myArgs := flag.Args()

	if invalidFlags(*deleteFlag, *deleteRFlag, *selectIndex) {
		fmt.Println("Invalid flags: too many flags given")
		return
	}

	// myIndex is the index given to strCopy function, defaults to 0 but can be changed
	// with selectIndex flag. Ex: subs -i 2
	myIndex := 0

	// directories is the list of directories given as arguments
	var directories []string

	// if an index was specified with the selectIndex flag, then it should be the
	// first item in myArgs variable
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
		// if directories is empty then use the current dir
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

// invalidFlags checks if more than one flag were given as arguments,
// which is not valid
func invalidFlags(delete bool, deleteR bool, index bool) bool {
	flagCount := 0
	if delete {
		flagCount++
	}
	if deleteR {
		flagCount++
	}
	if index {
		flagCount++
	}
	return flagCount > 1
}

// applyFlags checks which flag was given and calls the correct function
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
// 1) change copy function not to rely on bash
