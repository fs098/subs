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

func srtCopy(dir string) {
	var videoInfos []videoInfo
	getVideoInfo(dir, &videoInfos)
	subs := findSubsFileAndGetSubtitles(dir)

	if len(subs) == 0 {
		fmt.Println("No .srt files found")
		return
	}

	var counter int
	if len(subs) == 1 {
		copy(videoInfos[0].subtitleName, subs[0].subFiles[0])
		counter++
		fmt.Println("Copied", counter, "files!")
		return
	}

	for i, info := range videoInfos {
		if info.name == subs[i].folderName {
			copy(info.subtitleName, subs[i].subFiles[0])
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
