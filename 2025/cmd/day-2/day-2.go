package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/goforj/godump"
)

type ProductId struct {
	start int
	end   int
}

func getProductIds() []ProductId {
	file, err := os.ReadFile("./day-2.input")

	if err != nil {
		panic(err)
	}

	str := strings.Replace(string(file), "\n", "", 1)

	ids := []ProductId{}

	lines := strings.Split(str, ",")

	for _, line := range lines {
		p := ProductId{}
		l := strings.Split(line, "-")
		if len(l) != 2 {
			continue
		}

		p.start, _ = strconv.Atoi(l[0])
		p.end, _ = strconv.Atoi(l[1])

		ids = append(ids, p)
	}

	return ids
}

func IsInvalidId(id string) bool {
	str_len := len(id)

	if str_len == 2 {
		if id[0] != id[1] {
			return false
		}
	}

	for i := range str_len {
		str := id[0 : i+1]
		count := strings.Count(id, str)

		if count == str_len {
			return true
		}

		if count == 1 {
			return false
		}

		if len(str) == 1 {
			continue
		}

		if len(str)*count == str_len {
			return true
		}

	}
	return false
}

func main() {
	c := 0

	ids := getProductIds()

	for _, x := range ids {
		for i := x.start; i <= x.end; i++ {
			str := strconv.Itoa(i)
			middle := len(str) / 2
			z := str[0:middle]
			x := str[middle:]
			if z == x {
				c = c + i
			}
		}
	}
	godump.Dump("Day 2 part 1", c)

	c = 0

	for _, x := range ids {
		for i := x.start; i <= x.end; i++ {
			if i < 10 {
				continue
			}
			str := strconv.Itoa(i)
			isInvalid := IsInvalidId(str)
			if isInvalid {
				c += i
			}
		}
	}

	godump.Dump("Day 2 part 2", c)
}
