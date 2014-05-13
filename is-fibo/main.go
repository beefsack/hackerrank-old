package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
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
		num, _ := strconv.ParseFloat(l, 64)
		output := "IsNotFibo"
		if isFib(num) {
			output = "IsFibo"
		}
		outputs = append(outputs, output)
	}
	return strings.Join(outputs, "\n")
}

func isFib(n float64) bool {
	a := 5 * math.Pow(n, 2)
	return isWhole(math.Sqrt(a+4)) || isWhole(math.Sqrt(a-4))
}

func isWhole(n float64) bool {
	return math.Abs(float64(int(n+0.5))-n) < 0.000001
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
