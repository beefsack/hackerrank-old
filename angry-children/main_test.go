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
3
10
100
300
200
1000
20
30`,
			Expected: `20`,
		},
		Case{
			Input: `10
4
1
2
3
4
10
20
30
40
100
200`,
			Expected: `3`,
		},
	} {
		actual := run(bytes.NewBufferString(c.Input))
		if actual != c.Expected {
			t.Fatal(actual)
		}
	}
}
