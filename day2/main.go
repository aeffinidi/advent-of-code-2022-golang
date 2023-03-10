package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	// used in part 1
	win  = "W"
	loss = "L"
	draw = "D"
	// used in part 2
	rock     = "A"
	paper    = "B"
	scissors = "C"
)

func GetInput(path string) [][]string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	lines := strings.Split(string(b), "\n")
	var input [][]string = make([][]string, len(lines))
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		lineSplit := strings.Split(l, " ")
		set := []string{
			lineSplit[0],
			lineSplit[1],
		}
		input[i] = set

	}

	return input
}

func CounterAttack(choice string) string {
	var counterAttackStrategy = map[string]string{
		"A": "Y",
		"B": "X",
		"C": "Z",
	}

	if v, ok := counterAttackStrategy[choice]; ok {
		return v
	}

	return ""
}

func CounterAttack2(choice, outcome string) string {
	if outcome == "Y" {
		return choice
	}
	decisionMap := map[string]map[string]string{
		"X": {rock: scissors, paper: rock, scissors: paper},
		"Z": {rock: paper, paper: scissors, scissors: rock}}

	return decisionMap[outcome][choice]
}

func PlayResult(opponentChoice, myChoice string) string {
	switch opponentChoice {
	case "A":
		if myChoice == "Y" {
			return win
		}
		if myChoice == "Z" {
			return loss
		}
	case "B":
		if myChoice == "Z" {
			return win
		}
		if myChoice == "X" {
			return loss
		}
	case "C":
		if myChoice == "X" {
			return win
		}
		if myChoice == "Y" {
			return loss
		}
	}

	return "D"
}

func CalculateScore(result, myChoice string) int {
	// resultScore := map[string]int{win: 6, loss: 0, draw: 3}
	// choiceScore := map[string]int{"X": 1, "Y": 2, "Z": 3}
	resultScore := map[string]int{"Z": 6, "X": 0, "Y": 3}
	choiceScore := map[string]int{rock: 1, paper: 2, scissors: 3}
	return resultScore[result] + choiceScore[myChoice]
}

func SumScores2(rounds [][]string) int {
	sum := 0
	for _, v := range rounds {
		if len(v) == 0 {
			continue
		}
		score := CalculateScore(v[1], CounterAttack2(v[0], v[1]))
		sum += score
	}
	return sum
}

func SumScores(rounds [][]string) int {
	sum := 0
	for _, v := range rounds {
		if len(v) == 0 {
			continue
		}
		score := CalculateScore(PlayResult(v[0], v[1]), v[1])
		sum += score
	}
	return sum
}

func main() {
	input := GetInput("./input")
	// sum := SumScores(input)
	sum2 := SumScores2(input)

	// fmt.Printf("day-2 part 1: %d", sum)
	fmt.Printf("day-2 part 2: %d", sum2)
}
