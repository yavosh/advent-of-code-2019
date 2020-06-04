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
		result := getFuelForMass(testCase.weight)
		assert.Equal(t, testCase.expectedFuel, result, fmt.Sprintf("for weight %d [%d!=%d] ", testCase.weight, testCase.expectedFuel, result))
	}

}

func TestGetFuelIncludingFuel(t *testing.T) {
	testCases := []struct {
		weight            int
		expectedFuel      int
		expectedFuelFuel  int
		expectedTotalFuel int
	}{
		{1969, 654, 312, 966},
		{100756, 33583, 16763, 50346},
		{140916, 46970, 23455, 70425},
	}

	for _, testCase := range testCases {
		result := getFuelForMass(testCase.weight)
		resultFuel := getFuelMassFuel(testCase.weight)
		assert.Equal(t, testCase.expectedFuel, result,
			fmt.Sprintf("for weight %d fuel [%d!=%d] ", testCase.weight, testCase.expectedFuel, result))
		assert.Equal(t, testCase.expectedFuelFuel, resultFuel,
			fmt.Sprintf("for weight %d fuel for fuel requirement [%d!=%d] ", testCase.weight, testCase.expectedFuelFuel, resultFuel))
		assert.Equal(t, testCase.expectedTotalFuel, result+resultFuel,
			fmt.Sprintf("for weight %d total fuel requirement [%d!=%d] ", testCase.weight, testCase.expectedTotalFuel, result+resultFuel))
	}
}
