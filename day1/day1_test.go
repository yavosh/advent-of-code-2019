package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFuel(t *testing.T) {
	testCases := []struct {
		weight       int
		expectedFuel int
	}{
		{12, 2},
		{14, 2},
		{140916, 46970},
		{1969, 654},
		{100756, 33583},
	}

	for _, testCase := range testCases {
		result := getFuel(testCase.weight)
		assert.Equal(t, testCase.expectedFuel, result, fmt.Sprintf("for weight %d [%d!=%d] ", testCase.weight, testCase.expectedFuel, result))
	}

}
