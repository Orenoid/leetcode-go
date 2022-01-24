package main

import (
	"testing"
)

func Test344(t *testing.T) {
	s := []byte("abc")
	reverseString(s)
	if string(s) != "cba" {
		t.Error("wrong")
	}
}
