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
	outputs := []string{}
	b, _ := ioutil.ReadAll(input)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	testCases := parseInts(lines[0])[1]
	widths := parseInts(lines[1])
	for i := 2; i < testCases+2; i++ {
		ints := parseInts(lines[i])
		minWidth := -1
		for j := ints[0]; j <= ints[1]; j++ {
			if minWidth == -1 || widths[j] < minWidth {
				minWidth = widths[j]
			}
		}
		outputs = append(outputs, strconv.Itoa(minWidth))
	}
	return strings.Join(outputs, "\n")
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
