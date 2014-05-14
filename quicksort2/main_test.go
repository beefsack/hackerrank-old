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
			Input: `7
5 8 1 3 7 9 2`,
			Expected: `2 3
1 2 3
7 8 9
1 2 3 5 7 8 9`,
		},
	} {
		actual := run(bytes.NewBufferString(c.Input))
		if actual != c.Expected {
			t.Fatal(actual)
		}
	}
}
