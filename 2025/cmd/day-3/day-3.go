package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/goforj/godump"
)

func getLines() [][]int {
	file, err := os.ReadFile("./day-3.sample")

	if err != nil {
		panic(err)
	}

	ids := [][]int{}

	for line := range bytes.SplitSeq(file, []byte("\n")) {
		z := []int{}
		for n := range line {
			i, err := strconv.Atoi(string(line[n]))
			if err != nil {
				panic(err)
			}
			z = append(z, i)
		}

		if len(z) > 0 {
			ids = append(ids, z)
		}
	}

	return ids
}

func main() {
	lines := getLines()

	total := 0

	for _, line := range lines {
		tmp := 0
		for x := range line {
			left := line[x]
			right := line[x+1:]
			if len(right) == 0 {
				continue
			}

			highest := slices.Max(right)
			number := fmt.Sprintf("%d%d", left, highest)
			total, err := strconv.Atoi(number)

			if err != nil {
				panic(err)
			}
			if total > tmp {
				tmp = total
			}
		}
		total += tmp
	}

	godump.Dump("day 3 part 1", total)
}
