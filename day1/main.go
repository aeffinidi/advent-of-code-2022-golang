package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func ReadCaloriesFromInput(input string) (calories [][]int, err error) {
	lines := strings.Split(input, "\n")
	calories = make([][]int, 0)
	currentCalories := make([]int, 0)
	for _, l := range lines {
		if len(l) == 0 {
			calories = append(calories, currentCalories)
			currentCalories = make([]int, 0)
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		currentCalories = append(currentCalories, n)
	}
	if len(currentCalories) > 0 {
		calories = append(calories, currentCalories)
	}

	return calories, nil
}

func main() {
	b, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err.Error())
	}
	calories, e := ReadCaloriesFromInput(string(b))
	if e != nil {
		panic("Coundl't read from input file")
	}

	fmt.Printf("day-1 part 1: %d", MaxCaloryCarrier(SumCalories(calories)))
	fmt.Printf("day-1 part 2: %d", lo.Sum(MaxCalory3Carriers(SumCalories(calories))))
}

func MaxCaloryCarrier(calories []int) int {
	return lo.Max(calories)
}

func MaxCalory3Carriers(calories []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	return calories[:3]
}

func SumCalories(calories [][]int) []int {
	sums := make([]int, 0)
	for _, v := range calories {
		sums = append(sums, lo.Sum(v))
	}

	return sums
}
