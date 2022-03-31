package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	revisionSlice1 := strings.Split(version1, ".")
	revisionSlice2 := strings.Split(version2, ".")
	for i := 0; i < len(revisionSlice1) || i < len(revisionSlice2); i++ {
		revision1, revision2 := 0, 0
		if i < len(revisionSlice1) {
			revision1, _ = strconv.Atoi(revisionSlice1[i])
		}
		if i < len(revisionSlice2) {
			revision2, _ = strconv.Atoi(revisionSlice2[i])
		}
		if revision1 < revision2 {
			return -1
		}
		if revision2 < revision1 {
			return 1
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1.001", "1.02"))
}
