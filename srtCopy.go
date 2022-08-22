package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type videoInfo struct {
	name         string // without termination
	subtitleName string
	// path      string

}

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
		fmt.Println("Invalid index provided. Index is now", index)
	} else if index < 0 {
		index = 0
		fmt.Println("Invalid index provided. Index is now", index)
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

func getVideoInfo(dir string, infos *[]videoInfo) {
	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range dirContent {
		if file.IsDir() {
			getVideoInfo(dir+string(os.PathSeparator)+file.Name(), infos)
		}
		if isVideo(file.Name()) {
			newInfo := videoInfo{
				name:         getVideoName(file.Name()),
				subtitleName: dir + string(os.PathSeparator) + getVideoName(file.Name()) + ".srt",
				// path:      dir + string(os.PathSeparator) + file.Name(),

			}
			*infos = append(*infos, newInfo)
		}
	}
}

func isVideo(file string) bool {
	mkv := strings.HasSuffix(file, ".mkv")
	mp4 := strings.HasSuffix(file, ".mp4")

	return mkv || mp4
}

func getVideoName(name string) string {
	var vidTermination string

	if strings.HasSuffix(name, ".mkv") {
		vidTermination = ".mkv"
	} else if strings.HasSuffix(name, ".mp4") {
		vidTermination = ".mp4"
	}
	return strings.Replace(name, vidTermination, "", 1)
}
