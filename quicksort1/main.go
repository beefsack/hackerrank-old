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
	nums := parseInts(lines[1])
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
	for _, n := range l {
		outputs = append(outputs, strconv.Itoa(n))
	}
	outputs = append(outputs, strconv.Itoa(p))
	for _, n := range r {
		outputs = append(outputs, strconv.Itoa(n))
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
