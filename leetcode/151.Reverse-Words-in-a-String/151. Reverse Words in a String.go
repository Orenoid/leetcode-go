package main

import (
	"fmt"
)

type word struct {
	begin int
	end   int
}

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	sBytes := []byte(s)

	// 去除头尾的空格
	begin, end := 0, len(sBytes)-1
	for begin <= end && sBytes[begin] == ' ' {
		begin++
	}
	if begin > end {
		return ""
	}
	for end > begin && sBytes[end] == ' ' {
		end--
	}
	sBytes = sBytes[begin : end+1]
	// 去除单词之间多余的空格
	slow, fast := 0, 0
	for ; fast < len(sBytes); fast++ {
		if fast > 0 && sBytes[fast] == ' ' && sBytes[fast-1] == sBytes[fast] {
			continue
		}
		sBytes[slow] = sBytes[fast]
		slow++
	}
	sBytes = sBytes[:slow]
	//return string(sBytes)

	wordBegin := 0
	var words []word
	for i, ch := range sBytes {
		if ch == ' ' {
			words = append(words, word{begin: wordBegin, end: i - 1})
			wordBegin = i + 1
		} else if i == len(sBytes)-1 {
			words = append(words, word{begin: wordBegin, end: i})
		}
	}
	for _, word := range words {
		reverse(sBytes, word.begin, word.end)
	}
	reverse(sBytes, 0, len(sBytes)-1)
	return string(sBytes)
}

func reverse(sBytes []byte, begin int, end int) {
	for begin < end {
		sBytes[begin], sBytes[end] = sBytes[end], sBytes[begin]
		begin++
		end--
	}
}

func main() {
	fmt.Println(reverseWords("     the   sky  is blue   "))
}
