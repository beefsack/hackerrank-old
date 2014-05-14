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
	pieces := map[int][]string{}
	for k, l := range lines[1:] {
		parts := strings.Split(strings.TrimSpace(l), " ")
		i, _ := strconv.Atoi(parts[0])
		t := parts[1]
		if k < len(lines)/2 {
			t = "-"
		}
		if pieces[i] == nil {
			pieces[i] = []string{}
		}
		pieces[i] = append(pieces[i], t)
	}
	for i := 0; i < 100; i++ {
		if pieces[i] != nil {
			outputs = append(outputs, pieces[i]...)
		}
	}
	return strings.Join(outputs, " ")
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

func intstoa(ints []int, delim string) string {
	s := make([]string, len(ints))
	for k, i := range ints {
		s[k] = strconv.Itoa(i)
	}
	return strings.Join(s, delim)
}
