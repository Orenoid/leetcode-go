package main

import "fmt"

func exist(board [][]byte, word string) bool {

	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	found := false
	var backTrack func(i, j int, track []byte)
	backTrack = func(i, j int, track []byte) {
		if found || i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || used[i][j] {
			return
		}
		if word[len(track)] != board[i][j] {
			return
		}
		track = append(track, board[i][j])
		used[i][j] = true
		if len(track) == len(word) {
			found = true
		} else {
			backTrack(i-1, j, track)
			backTrack(i, j-1, track)
			backTrack(i+1, j, track)
			backTrack(i, j+1, track)
		}
		track = track[:len(track)-1]
		used[i][j] = false
	}
	for i := range board {
		for j := range board[0] {
			if !found && board[i][j] == word[0] {
				backTrack(i, j, nil)
			}
		}
	}
	return found
}

func main() {
	fmt.Println(exist([][]byte{{'a', 'b'}, {'c', 'd'}}, "bca"))
}
