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
	for _, l := range lines[1:] {
		ints := parseInts(l)
		n, c, m := ints[0], ints[1], ints[2]
		choc := n / c
		wrap := choc
		for wrap >= m {
			choc += 1
			wrap -= m - 1
		}
		outputs = append(outputs, strconv.Itoa(choc))
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
