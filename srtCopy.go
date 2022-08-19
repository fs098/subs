package main

import (
	"fmt"
	"os"
	"strings"
)

type videoInfo struct {
	name         string // without termination
	path         string
	subtitleName string
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
				path:         dir + string(os.PathSeparator) + file.Name(),
				subtitleName: dir + string(os.PathSeparator) + getVideoName(file.Name()) + ".srt",
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

func strCopyTest() {
	// cmd := exec.Command("cp", "oldDir", "newDir")
	// cmd.Run()

	// test1 := getVideoName("test1.mkv")
	// test2 := getVideoName("test2.mp4")
	// fmt.Println(test1, test2)

	flagsOrDir := os.Args[1:]
	myVids := []videoInfo{}
	getVideoInfo(flagsOrDir[0], &myVids)
	for _, vid := range myVids {
		fmt.Println("---------------------------")
		fmt.Println("name:", vid.name)
		fmt.Println("path:", vid.path)
		fmt.Println("sub name:", vid.subtitleName)
	}
}
