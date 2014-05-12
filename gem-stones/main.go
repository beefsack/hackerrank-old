package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(run(os.Stdin))
}

func run(input io.Reader) string {
	b, _ := ioutil.ReadAll(input)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	numLines, _ := strconv.Atoi(lines[0])
	numLinesIn := map[rune]int{}
	for _, l := range lines[1:] {
		inLine := map[rune]bool{}
		for _, r := range l {
			inLine[r] = true
		}
		for r, _ := range inLine {
			numLinesIn[r] += 1
		}
		// Logic goes here
	}
	gems := 0
	for _, n := range numLinesIn {
		if n == numLines {
			gems += 1
		}
	}
	return strconv.Itoa(gems)
}

func parseInts(input string) []int {
	strInts := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(input), -1)
	ints := make([]int, len(strInts))
	for k, s := range strInts {
		i, _ := strconv.Atoi(s)
		ints[k] = i
	}
	return ints
}
