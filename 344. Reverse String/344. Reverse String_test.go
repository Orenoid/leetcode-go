package main

import "testing"

func Test(t *testing.T) {
	s := []byte("abc")
	reverseString(s)

	if string(s) != "cba" {
		t.Error("wrong")
	}
}
