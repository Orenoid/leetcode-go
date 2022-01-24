package main

import "testing"
import "github.com/stretchr/testify/assert"

type question struct {
	param
	ret
}

type param struct {
	one string
}

type ret struct {
	one string
}

func TestOffer5(t *testing.T) {
	questions := []question{
		{
			param{"hello world"},
			ret{"hello%20world"},
		},
		{
			param{" hello"},
			ret{"%20hello"},
		},
		{
			param{"  "},
			ret{"%20%20"},
		},
		{
			param{"hello"},
			ret{"hello"},
		},
		{
			param{"hello "},
			ret{"hello%20"},
		},
	}
	for _, q := range questions {
		assert.Equal(t, q.ret.one, replaceSpace(q.param.one))
	}
}
