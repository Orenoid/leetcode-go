package main

import "fmt"

func exist(board [][]byte, word string) bool {
	found := false
	usedFlags := make([][]bool, len(board))
	for i := range usedFlags {
		usedFlags[i] = make([]bool, len(board[0]))
	}

	var backTrack func(i, j, idx int)
	backTrack = func(i, j, idx int) {
		if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 ||
			idx > len(word)-1 || usedFlags[i][j] || board[i][j] != word[idx] || found {
			return
		}
		usedFlags[i][j] = true
		if idx == len(word)-1 {
			found = true
		} else {
			backTrack(i-1, j, idx+1)
			backTrack(i, j-1, idx+1)
			backTrack(i+1, j, idx+1)
			backTrack(i, j+1, idx+1)
		}
		usedFlags[i][j] = false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backTrack(i, j, 0)
		}
	}
	return found
}

func main() {
	fmt.Println(exist([][]byte{{'a', 'b'}, {'c', 'd'}}, "bdc"))
}
