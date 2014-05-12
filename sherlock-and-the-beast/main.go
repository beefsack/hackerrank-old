package main

import (
	"bytes"
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
		digits, _ := strconv.Atoi(l)
		key := &bytes.Buffer{}
		for digits >= 3 {
			if digits < 15 && digits%5 == 0 {
				key.WriteString(strings.Repeat("3", digits))
				digits = 0
			} else {
				key.WriteString("555")
				digits -= 3
			}
		}
		output := "-1"
		if digits == 0 {
			output = key.String()
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
