package main

import (
	"fmt"
	"os"
	"os/exec"
)

func srtCopy(dir string, index int) {
	var videoInfos []videoInfo
	getVideoInfo(dir, &videoInfos)
	subs := findSubsFileAndGetSubtitles(dir)

	if len(subs) == 0 {
		fmt.Println("No .srt files found")
		return
	}

	// in case the index given isn't valid
	maxIndex := len(subs[0].subFiles) - 1
	if index > maxIndex {
		index = maxIndex
		fmt.Println("Invalid index provided. Index is now", index+1)
	} else if index < 0 {
		index = 0
		fmt.Println("Invalid index provided. Index is now", index+1)
	}

	var counter int
	// Single video file
	if len(subs) == 1 {
		copy(videoInfos[0].subtitleName, subs[0].subFiles[index])
		counter++
		fmt.Println("Copied", counter, "files!")
		return
	}

	// Multiple videos
	for i, info := range videoInfos {
		if info.name == subs[i].folderName {
			copy(info.subtitleName, subs[i].subFiles[index])
			counter++
		}
	}
	fmt.Println("Copied", counter, "files!")
}

func copy(destFile string, srcFile string) {
	cmd := exec.Command("cp", srcFile, destFile)
	cmd.Run()
}

func getPath(dir string, name string) string {
	if dir[len(dir)-1] == '/' {
		return dir + name
	}
	return dir + string(os.PathSeparator) + name
}
