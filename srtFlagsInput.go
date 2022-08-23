package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// applyFlagsHandler verifies if given dir has a subs folder and calls applyFlags function.
// If the folder is not found then searches the directory for folders with a subs folder
// and appends them to a list. Then gets input from the user as to which folders from
// the list the program should use (with getInput function), and calls applyFlags function
// to each folder.
func applyFlagsHandler(dir string, index int, delete bool, deleteR bool) {
	if hasSubs(dir) {
		applyFlags(dir, index, delete, deleteR)
		return
	}

	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	var foldersWithSubs []string
	for _, file := range dirContent {
		if !file.IsDir() {
			continue
		}
		folderName := getPath(dir, file.Name())
		if hasSubs(folderName) {
			foldersWithSubs = append(foldersWithSubs, folderName)
		}
	}

	if len(foldersWithSubs) == 0 {
		fmt.Println("No valid folders found in this directory")
		return
	}

	myInts, err := getInput(foldersWithSubs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, n := range myInts {
		applyFlags(foldersWithSubs[n], index, delete, deleteR)
	}
}

// getInput lists the folders given as argument and it's indexes, starting at 1,
// and gets user input as to which of them the program should use, returns
// inputed indexes in a list or an error.
func getInput(folders []string) ([]int, error) {
	var myInts []int

	fmt.Println("Found the following folders:")
	for i, folder := range folders {
		fmt.Printf("[%d] - %s\n", i+1, folder)
	}

	fmt.Println("Select the folders to appy subs. Default: All")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return myInts, err
	}
	input := strings.Replace(line, "\n", "", 1)

	// Default case:
	if input == "" {
		for i := range folders {
			myInts = append(myInts, i)
		}
		return myInts, nil
	}

	// get indexes given as input
	myIntsString := strings.Split(input, " ")
	for _, s := range myIntsString {
		myInt, err := strconv.Atoi(s)
		if err != nil {
			return myInts, err
		}
		myInts = append(myInts, myInt-1)
	}
	return myInts, nil
}

// hasSubs checks if given directory has a folder with the name "Subs"
func hasSubs(dir string) bool {
	_, err := os.Stat(getPath(dir, "Subs"))
	if !os.IsNotExist(err) {
		return true
	}
	_, err = os.Stat(getPath(dir, "subs"))
	return !os.IsNotExist(err)
}
