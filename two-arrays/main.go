package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
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
	testCases, _ := strconv.Atoi(lines[0])
	for i := 0; i < testCases; i++ {
		output := "YES"
		k := parseInts(lines[i*3+1])[1]
		arr1 := parseInts(lines[i*3+2])
		sort.Ints(arr1)
		arr2 := parseInts(lines[i*3+3])
		sort.Ints(arr2)
		for j := 0; j < len(arr1); j++ {
			if arr1[j]+arr2[len(arr2)-j-1] < k {
				output = "NO"
				break
			}
		}
		outputs = append(outputs, output)
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

func intstoa(ints []int, delim string) string {
	s := make([]string, len(ints))
	for k, i := range ints {
		s[k] = strconv.Itoa(i)
	}
	return strings.Join(s, delim)
}
