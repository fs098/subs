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
			result = srtFetcher(getPath(dir, file.Name()))
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
		fileSubs := getSubFiles(getPath(dir, file.Name()))
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
			result = append(result, getPath(dir, file.Name()))
		}
	}

	return result
}
