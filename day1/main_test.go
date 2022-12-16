package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxCaloryCarrier(t *testing.T) {
	assert.Equal(t, 24000, MaxCaloryCarrier([]int{6000, 4000, 11000, 24000, 10000}))
}

func Test_MaxCalory3Carriers(t *testing.T) {
	assert.Equal(t, []int{24000, 11000, 10000}, MaxCalory3Carriers([]int{6000, 4000, 11000, 24000, 10000}))
}

func Test_SumCalories(t *testing.T) {
	assert.Equal(t, []int{2}, SumCalories([][]int{{1, 1}}))
	assert.Equal(t, []int{2, 3}, SumCalories([][]int{{1, 1}, {1, 2}}))
	assert.Equal(t, []int{1, 2, 3, 4}, [][]int{{1}, {1, 1}, {1, 1, 1}, {1, 1, 1, 1}})
}

func Test_ReadInput_OneLineInput(t *testing.T) {
	output, _ := ReadCaloriesFromInput("1000")
	assert.Equal(t, [][]int{{1000}}, output)
}

func Test_ReadInput(t *testing.T) {

	b, err := ioutil.ReadFile("./input_sample")
	if err != nil {
		panic(err.Error())
	}
	output, _ := ReadCaloriesFromInput(string(b))

	assert.Equal(t, [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}, output)
}
