package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindSameItem(t *testing.T) {
	assert.Equal(t, 'p', FindSameItem("vJrwpWtwJgWr", "hcsFMMfFFhFp"))
	assert.Equal(t, 'L', FindSameItem("jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"))
	assert.Equal(t, 'P', FindSameItem("PmmdzqPrV", "vPwwTWBwg"))
	assert.Equal(t, int32(0), FindSameItem("mDSlGBGwhGhm", "LHCQnMMVMLMVFJ"))
}

type output struct {
	c1 string
	c2 string
}

func Test_DivideLine(t *testing.T) {
	tests := []struct {
		input  string
		output output
	}{
		{input: "DlLqDPwGlbVCdVbddbsr",
			output: output{
				c1: "DlLqDPwGlb", c2: "VCdVbddbsr",
			}},
		{input: "qgFtChDCFgmCnppPnczPbcLpnL",
			output: output{
				c1: "qgFtChDCFgmCn", c2: "ppPnczPbcLpnL",
			}},
		{input: "ZrMWTTbbGMpbtgCTTZgZCWCLhwdsLvhhhPhNSldvPwPNfNlR",
			output: output{
				c1: "ZrMWTTbbGMpbtgCTTZgZCWCL", c2: "hwdsLvhhhPhNSldvPwPNfNlR",
			}},
		{input: "DDZlblRRcNpJNhpL",
			output: output{
				c1: "DDZlblRR", c2: "cNpJNhpL",
			}},
		{input: "rPCMQnbdLqJLzhzB",
			output: output{
				c1: "rPCMQnbd", c2: "LqJLzhzB",
			}},
		{input: "cDczzSMzDcSGSQbFjRFZtZCZmtwZRt",
			output: output{
				c1: "cDczzSMzDcSGSQb", c2: "FjRFZtZCZmtwZRt",
			}},
		{input: "bVcJSbVbSDNrrWWNdvWf",
			output: output{
				c1: "bVcJSbVbSD", c2: "NrrWWNdvWf",
			}},
		{input: "LzLrzDhmDRRJpJzptDhCqHMHqnqSlHqMSQNHQS",
			output: output{
				c1: "LzLrzDhmDRRJpJzptDh", c2: "CqHMHqnqSlHqMSQNHQS",
			}},
		{input: "dVRRTCTVLcfJcJFb",
			output: output{
				c1: "dVRRTCTV", c2: "LcfJcJFb",
			}},
		{input: "RRjgNPTRFhglgNNjTsqCCGZfzmHCnZGnZCqq",
			output: output{
				c1: "RRjgNPTRFhglgNNjTs", c2: "qCCGZfzmHCnZGnZCqq",
			}},
		{input: "dSDbbJdlHFNlll",
			output: output{
				c1: "dSDbbJd", c2: "lHFNlll",
			}},
	}
	for _, test := range tests {
		c1, c2 := DivideLine(test.input)
		assert.Equal(t, test.output.c1, c1)
		assert.Equal(t, test.output.c2, c2)
	}
}

func Test_CalculateSameItem(t *testing.T) {
	assert.Equal(t, 0, CalculateSameItem([]string{""}))

	assert.Equal(t, 96, CalculateSameItem([]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}))

	assert.Equal(t, 61, CalculateSameItem([]string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}))

	assert.Equal(t, 157, CalculateSameItem([]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}))
}

func Test_ItemScore(t *testing.T) {
	// lowercase
	assert.Equal(t, 16, ItemScore('p'))
	assert.Equal(t, 1, ItemScore('a'))

	// uppercase
	assert.Equal(t, 38, ItemScore('L'))
	assert.Equal(t, 42, ItemScore('P'))
}
