package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPass(t *testing.T) {

	testCases := []struct {
		pass   string
		valid  bool
		reason string
	}{
		{"a11111", false, "format"},
		{"1111", false, "length"},
		{"111111", false, "no-double"},
		{"123455", true, "ok"},
		{"223450", false, "decreasing"},
		{"123789", false, "no-double"},
		{"112233", true, "ok"},
		{"123444", false, "no-double"},
		{"111122", true, "ok"},
		{"111222", false, "no-double"},
	}

	for _, testCase := range testCases {
		result, reason := validPass(testCase.pass)
		assert.Equal(t, testCase.valid, result, fmt.Sprintf("expecting %q to [%t!=%t]", testCase.pass, testCase.valid, result))
		assert.Equal(t, testCase.reason, reason, fmt.Sprintf("expecting %q to have reason [%s!=%s]", testCase.pass, testCase.reason, reason))
	}
}

func TestValidSinglePass(t *testing.T) {
	result, reason := validPass("111222")
	assert.Equal(t, false, result)
	assert.Equal(t, "no-double", reason)
}
