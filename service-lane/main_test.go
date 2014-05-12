package main

import (
	"bytes"
	"testing"
)

type Case struct {
	Input, Expected string
}

func TestRun(t *testing.T) {
	for _, c := range []Case{
		Case{
			Input: `8 5
2 3 1 2 3 2 3 3
0 2
0 1
6 7
3 5
0 7`,
			Expected: `1
2
3
2
1`,
		},
	} {
		actual := run(bytes.NewBufferString(c.Input))
		if actual != c.Expected {
			t.Fatal(actual)
		}
	}
}
