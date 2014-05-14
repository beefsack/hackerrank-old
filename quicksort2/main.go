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
	_, outputs := sort(nums)
	return strings.Join(outputs, "\n")
}

func sort(nums []int) (res []int, outputs []string) {
	if len(nums) <= 1 {
		res = nums
		return
	}
	p := nums[0]
	l := []int{}
	r := []int{}
	for _, n := range nums[1:] {
		if n < p {
			l = append(l, n)
		} else {
			r = append(r, n)
		}
	}
	res, lOutputs := sort(l)
	outputs = lOutputs
	res = append(res, p)
	r, rOutputs := sort(r)
	outputs = append(outputs, rOutputs...)
	res = append(res, r...)
	outputs = append(outputs, intstoa(res, " "))
	return
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
