package main

import (
	"fmt"
	"os"
	"strings"
)

func getSubtitles(dir string) []string {
	var result []string

	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range dirContent {
		if strings.HasSuffix(file.Name(), ".srt") && strings.Contains(strings.ToLower(file.Name()), "english") {
			result = append(result, dir+string(os.PathSeparator)+file.Name())
		}
	}

	return result
}

func srtFetcherTest() {
	flagsOrDir := os.Args[1:]
	mySubs := getSubtitles(flagsOrDir[0])
	fmt.Println(mySubs)
}
