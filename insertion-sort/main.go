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
	for tc := 0; tc < testCases; tc++ {
		nums := parseInts(lines[2*(tc+1)])
		sortedNums := parseInts(lines[2*(tc+1)])
		sort.Ints(sortedNums)
		shifts := 0
		for len(nums) > 0 {
			n := nums[0]
			// binary search
			lower := 0
			upper := len(sortedNums)
			k := 0
			for lower < upper {
				k = (lower + upper) / 2
				if sortedNums[k] == n {
					break
				} else if sortedNums[k] > n {
					upper = k
				} else {
					lower = k
				}
			}
			// make sure we get the first one
			for k > 0 && sortedNums[k-1] == n {
				k -= 1
			}
			shifts += k
			nums = nums[1:]
			sortedNums = append(sortedNums[:k], sortedNums[k+1:]...)
		}
		outputs = append(outputs, strconv.Itoa(shifts))
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
