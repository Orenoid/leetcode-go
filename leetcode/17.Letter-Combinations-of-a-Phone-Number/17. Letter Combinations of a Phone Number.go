package main

import "fmt"

var mapping = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var res []string
	var backTrack func(curr int, track []byte)
	backTrack = func(curr int, track []byte) {
		for _, letter := range mapping[digits[curr]] {
			track = append(track, letter)
			if len(track) == len(digits) {
				res = append(res, string(track))
			} else {
				backTrack(curr+1, track)
			}
			track = track[:len(track)-1]
		}
	}
	backTrack(0, nil)
	return res
}

func main() {
	fmt.Println(letterCombinations("2"))
}
