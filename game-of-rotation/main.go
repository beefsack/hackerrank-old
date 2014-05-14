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
	numLen := len(nums)

	var pmean, sum int
	// First pass, initial pmean and calc sum
	for k, n := range nums {
		pmean += (k + 1) * n
		sum += n
	}

	// Assuming each rotation reduces all by sum and increases the weight of
	// numLen
	highest := pmean
	for i := 0; i < numLen; i++ {
		pmean += sum - nums[numLen-i-1]*numLen
		if pmean > highest {
			highest = pmean
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
