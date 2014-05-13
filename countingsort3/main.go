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
	outputs := make([]string, 100)
	b, _ := ioutil.ReadAll(input)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	counts := map[int]int{}
	for _, l := range lines[1:] {
		parts := strings.Split(l, " ")
		n, _ := strconv.Atoi(parts[0])
		counts[n] += 1
	}
	runningCount := 0
	for i := 0; i < 100; i++ {
		runningCount += counts[i]
		outputs[i] = strconv.Itoa(runningCount)
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
