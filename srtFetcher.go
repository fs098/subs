package main

import (
	"fmt"
	"os"
	"strings"
)

type subtitles struct {
	folderName string
	subFiles   []string
}

func findSubsFileAndGetSubtitles(dir string) []subtitles {
	var result []subtitles

	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range dirContent {
		if !file.IsDir() {
			continue
		}
		if strings.ToLower(file.Name()) == "subs" {
			result = srtFetcher(dir + string(os.PathSeparator) + file.Name())
		}
	}
	return result
}

// srtFetcher fetches the subtitles from a subs directory,
// there are two cases:
// 1) if there is a single video file, all subtitles should be in a directory called subs
// 2) if there are multiple video files, there will be a folder of subtitles for each video
// inside the subs folder
func srtFetcher(dir string) []subtitles {
	// first case
	dirSubs := getSubFiles(dir)
	if len(dirSubs) > 0 {
		dirSubtitles := make([]subtitles, 1)
		dirSubtitles[0] = subtitles{
			folderName: "Subs",
			subFiles:   dirSubs,
		}
		return dirSubtitles
	}

	// second case
	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	var folderSubtitles []subtitles
	for _, file := range dirContent {
		if !file.IsDir() {
			continue
		}
		fileSubs := getSubFiles(dir + string(os.PathSeparator) + file.Name())
		if len(fileSubs) > 0 {
			fileSubtitles := subtitles{
				folderName: file.Name(),
				subFiles:   fileSubs,
			}
			folderSubtitles = append(folderSubtitles, fileSubtitles)
		}
	}
	return folderSubtitles
}

func getSubFiles(dir string) []string {
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
	// flagsOrDir := os.Args[1:]
	// mySubs := getSubFiles(flagsOrDir[0])
	// fmt.Println(mySubs)

	// flagsOrDir := os.Args[1:]
	// mySubs := srtFetcher(flagsOrDir[0])
	// for _, sub := range mySubs {
	// 	fmt.Println("---------------------------")
	// 	fmt.Println("sub folder:", sub.folderName)
	// 	fmt.Println("Subtitles:")
	// 	for i, subFile := range sub.subFiles {
	// 		fmt.Println(i, "=", subFile)
	// 	}
	// }

	flagsOrDir := os.Args[1:]
	mySubs := findSubsFileAndGetSubtitles(flagsOrDir[0])
	for _, sub := range mySubs {
		fmt.Println("---------------------------")
		fmt.Println("sub folder:", sub.folderName)
		fmt.Println("Subtitles:")
		for i, subFile := range sub.subFiles {
			fmt.Println(i, "=", subFile)
		}
	}
}
