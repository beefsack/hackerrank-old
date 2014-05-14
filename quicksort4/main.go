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
	nums1 := parseInts(lines[1])
	nums2 := parseInts(lines[1])
	return strconv.Itoa(insertionSortOps(nums1) - quickSortOps(nums2))
}

func insertionSortOps(nums []int) int {
	ops := 0
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				ops += 1
			}
		}
	}
	return ops
}

func quickSortOps(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	ops := 0
	p := nums[len(nums)-1]
	swap := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= p {
			nums[i], nums[swap] = nums[swap], nums[i]
			ops += 1
			swap += 1
		}
	}
	swap -= 1
	return ops + quickSortOps(nums[:swap]) + quickSortOps(nums[swap+1:])
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
