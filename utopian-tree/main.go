package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
	num, _ := strconv.Atoi(lines[0])
	for i := 1; i <= num; i++ {
		cycles, _ := strconv.Atoi(lines[i])
		h := 1
		for j := 0; j < cycles; j++ {
			if j%2 == 0 {
				h *= 2
			} else {
				h += 1
			}
		}
		outputs = append(outputs, strconv.Itoa(h))
	}
	return strings.Join(outputs, "\n")
}
