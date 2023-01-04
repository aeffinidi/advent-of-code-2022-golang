package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetInput(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	lines := strings.Split(string(b), "\n")
	var input []string = make([]string, len(lines))
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		input[i] = l
	}

	return input
}

func FindSameItem(line1 string, line2 string) rune {
	for _, c := range line1 {
		if strings.ContainsRune(line2, c) {
			return c
		}
	}

	// fmt.Printf("%s %s\n", line1, line2)
	return 0
}

func DivideLine(line string) (string, string) {
	length := len(line)
	c1, c2 := line[:length/2], line[(length)/2:]
	return c1, c2
}

func CalculateSameItem(lines []string) int {
	sum := 0
	for _, line := range lines {
		length := len(line)
		if length == 0 {
			continue
		}
		c1, c2 := DivideLine(line)
		item := FindSameItem(c1, c2)
		sum += ItemScore(item)
	}

	return sum
}

func ItemScore(ch rune) int {
	if int(ch) == 0 {
		return 0
	}
	offset := 96
	if ch <= 90 && ch >= 65 {
		offset = 38 // removed 27 from offset
	}

	return int(ch) - offset
}

func main() {
	input := GetInput("./input")
	sum1 := CalculateSameItem(input)

	fmt.Printf("day-3 part 1: %d\n", sum1)
	fmt.Printf("zero sample day-3 part 1: %d\n", CalculateSameItem(GetInput("./zero_sample")))
}
