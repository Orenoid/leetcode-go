package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	scanned := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		scanned[i] = make([]bool, len(grid[0]))
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if scanned[i][j] || grid[i][j] == '0' {
				continue
			}
			count++
			queue := [][2]int{{i, j}}
			for len(queue) > 0 {

				node := queue[0]
				queue = queue[1:]
				row, col := node[0], node[1]
				if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && !scanned[row][col] && grid[row][col] == '1' {
					scanned[row][col] = true
					queue = append(queue, [2]int{row - 1, col}, [2]int{row + 1, col}, [2]int{row, col - 1}, [2]int{row, col + 1})
				}

			}
		}
	}
	return count
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '0', '1'},
		{'0', '1', '0'},
	}))
}
