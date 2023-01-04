package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CounterAttack2(t *testing.T) {
	// opponent plays: Rock
	assert.Equal(t, rock, CounterAttack2(rock, "Y"))
	assert.Equal(t, scissors, CounterAttack2(rock, "X"))
	assert.Equal(t, paper, CounterAttack2(rock, "Z"))

	// opponent plays: Paper
	assert.Equal(t, paper, CounterAttack2(paper, "Y"))
	assert.Equal(t, rock, CounterAttack2(paper, "X"))
	assert.Equal(t, scissors, CounterAttack2(paper, "Z"))
}

func Test_SumScores2_Sample(t *testing.T) {
	assert.Equal(t, 4, SumScores2([][]string{{rock, "Y"}}))
	assert.Equal(t, 1, SumScores2([][]string{{paper, "X"}}))
	assert.Equal(t, 12, SumScores2([][]string{{rock, "Y"}, {paper, "X"}, {scissors, "Z"}}))
	assert.Equal(t, 19, SumScores2([][]string{{scissors, "X"}, {scissors, "X"}, {scissors, "X"}, {rock, "Z"}, {paper, "Y"}}))
}
