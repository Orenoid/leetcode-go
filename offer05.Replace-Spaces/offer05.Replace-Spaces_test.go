package main

import "testing"

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
		ret := replaceSpace(q.param.one)
		if ret != q.ret.one {
			t.Errorf("want: %v, got: %v", q.ret.one, ret)
		}
	}
}
