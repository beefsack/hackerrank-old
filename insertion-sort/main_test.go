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
			Input: `2
5
1 1 1 2 2
5
2 1 3 1 2`,
			Expected: `0
4`,
		},
		Case{
			Input:    monsterInput,
			Expected: monsterExpected,
		},
	} {
		actual := run(bytes.NewBufferString(c.Input))
		if actual != c.Expected {
			t.Fatal(actual)
		}
	}
}
