package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	param
	answer
}

type param struct {
	one []string
}

type answer struct {
	one int
}

func Test150(t *testing.T) {
	qs := []question{
		{
			param{[]string{"1", "2", "+"}},
			answer{3},
		},
		{
			param{[]string{"18"}},
			answer{18},
		},
		{
			param{[]string{"4", "13", "5", "/", "+"}},
			answer{6},
		},
	}
	for _, q := range qs {
		assert.Equal(t, q.answer.one, evalRPN(q.param.one))
	}
}
