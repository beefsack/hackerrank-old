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
	b, _ := ioutil.ReadAll(input)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	numPacks, _ := strconv.Atoi(lines[1])
	packs := []int{}
	for _, l := range lines[2:] {
		s, _ := strconv.Atoi(l)
		packs = append(packs, s)
	}
	sort.Ints(packs)
	minUnf := -1
	for k, i := range packs[:len(packs)-numPacks+1] {
		unf := packs[k+numPacks-1] - i
		if minUnf == -1 || unf < minUnf {
			minUnf = unf
		}
	}
	return strconv.Itoa(minUnf)
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
