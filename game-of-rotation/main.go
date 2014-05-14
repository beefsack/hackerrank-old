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
	nums := parseInts(lines[1])

	highest := 0
	for i := 0; i < len(nums); i++ {
		pm := pmean(nums, i, len(nums))
		if i == 0 || pm > highest {
			highest = pm
		}
	}
	return strconv.Itoa(highest)
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

var pmeanCache = map[int]map[int]int{}

func pmean(nums []int, start, length int) (mean int) {
	if length == 1 {
		return at(nums, start)
	}
	if pmeanCache[start] == nil {
		pmeanCache[start] = map[int]int{}
	}
	if _, ok := pmeanCache[start][length]; !ok {
		pmeanCache[start][length] =
			blockSum(nums, start, length) + pmean(nums, start+1, length-1)
	}
	return pmeanCache[start][length]
}

var blockSumCache = map[int]map[int]int{}

func blockSum(nums []int, start, length int) (sum int) {
	if length == 1 {
		return at(nums, start)
	}
	if blockSumCache[start] == nil {
		blockSumCache[start] = map[int]int{}
	}
	if _, ok := blockSumCache[start][length]; !ok {
		blockSumCache[start][length] =
			at(nums, start) + blockSum(nums, start+1, length-1)
	}
	return blockSumCache[start][length]
}

func at(nums []int, start int) int {
	return nums[start%len(nums)]
}
