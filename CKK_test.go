package CKK

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name     string
	inputArr []int
	res      Result
}

func TestCkkNode(t *testing.T) {
	tests := []testStruct{
		{
			name:     "Simple success test",
			inputArr: []int{4, 5, 6, 7, 8},
			res: Result{
				ResDiff:  0,
				ResLeft:  []int{7, 8},
				ResRight: []int{6, 5, 4},
			},
		},
		{
			name:     "Simple success test 2",
			inputArr: []int{8, 5, 7, 10, 6, 4},
			res: Result{
				ResDiff:  0,
				ResLeft:  []int{4, 6, 10},
				ResRight: []int{5, 7, 8},
			},
		},
		{
			name:     "Simple success test 3 - with negative numbers",
			inputArr: []int{8, 7, 6, 5, 4, -4, -4},
			res: Result{
				ResDiff:  0,
				ResLeft:  []int{-4, -4, 5, 6, 8},
				ResRight: []int{4, 7},
			},
		},
		{
			name:     "Success test (diff: 0)",
			inputArr: []int{528, 129, 376, 504, 543, 363, 213, 138, 206, 440, 504, 418},
			res: Result{
				ResDiff:  0,
				ResLeft:  []int{129, 138, 363, 504, 504, 543},
				ResRight: []int{206, 213, 376, 418, 440, 528},
			},
		},
		{
			name:     "Success test (diff: 1)",
			inputArr: []int{528, 129, 376, 504, 543, 363, 213, 138, 206, 440, 504, 419},
			res: Result{
				ResDiff:  1,
				ResLeft:  []int{129, 138, 213, 363, 376, 419, 543},
				ResRight: []int{206, 440, 504, 504, 528},
			},
		},
		{
			name:     "Success test (diff: > 1)",
			inputArr: []int{528, 129, 376, 504, 543, 363, 213, 138, 206, 440, 504, 424},
			res: Result{
				ResDiff:  2,
				ResLeft:  []int{129, 206, 363, 440, 504, 543},
				ResRight: []int{138, 213, 376, 424, 504, 528},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ckk, err := NewCkk(test.inputArr)
			assert.NoError(t, err)

			ckk.Run()
			res := ckk.GetResult()

			assert.Equal(t, res, test.res)
		})
	}
}
