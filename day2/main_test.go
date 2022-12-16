package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetInput(t *testing.T) {
	res := GetInput("./input_sample")
	assert.Len(t, res, 3)
	assert.Equalf(t, "Y", res[0][1], "for A")
	assert.Equalf(t, "X", res[1][1], "for B")
	assert.Equalf(t, "Z", res[2][1], "for C")
}

func Test_CounterAttack(t *testing.T) {
	assert.Equal(t, "Y", CounterAttack("A"))
	assert.Equal(t, "X", CounterAttack("B"))
	assert.Equal(t, "Z", CounterAttack("C"))
}

func Test_PlayResult_Win(t *testing.T) {
	assert.Equal(t, "W", PlayResult("A", "Y"))
	assert.Equal(t, "W", PlayResult("B", "Z"))
	assert.Equal(t, "W", PlayResult("C", "X"))
}

func Test_PlayResult_Loss(t *testing.T) {
	assert.Equal(t, "L", PlayResult("A", "Z"))
	assert.Equal(t, "L", PlayResult("B", "X"))
	assert.Equal(t, "L", PlayResult("C", "Y"))
}

func Test_PlayResult_Draw(t *testing.T) {
	assert.Equal(t, "D", PlayResult("A", "X"))
	assert.Equal(t, "D", PlayResult("B", "Y"))
	assert.Equal(t, "D", PlayResult("C", "Z"))
}

func Test_CalculateScore(t *testing.T) {
	assert.Equal(t, 8, CalculateScore(win, "Y"))
	assert.Equal(t, 1, CalculateScore(loss, "X"))
	assert.Equal(t, 6, CalculateScore(draw, "Z"))
}

func Test_SumScores_Sample(t *testing.T) {
	assert.Equal(t, 15, SumScores(GetInput("./input_sample")))
	assert.Equal(t, 29, SumScores(GetInput("./input_sample2")))
}
