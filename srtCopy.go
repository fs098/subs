package main

import "strings"

type videoInfo struct {
	name         string // without termination
	path         string
	subtitleName string
}

func isVideo(file string) bool {
	mkv := strings.HasSuffix(file, ".mkv")
	mp4 := strings.HasSuffix(file, ".mp4")

	return mkv || mp4
}

func strCopyTest() {
}
