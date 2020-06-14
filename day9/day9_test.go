package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yavosh/advent-of-code-2019/computer"
)

func TestNewCodeA(t *testing.T) {
	program := `109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`
	memory := computer.LoadInstructionsWithMemoryAlloc(program, 8192)
	_, out := computer.Run(memory, []int64{0})
	fmt.Printf("out=%v\n", out)
	assert.Equal(t, int64(99), out[0])
}

func TestNewCodeB(t *testing.T) {
	program := `1102,34915192,34915192,7,4,7,99,0`
	memory := computer.LoadInstructionsWithMemoryAlloc(program, 8192)
	_, out := computer.Run(memory, []int64{0})
	fmt.Printf("out=%v\n", out)
	assert.Equal(t, int64(1219070632396864), out[0])
}

func TestNewCodeC(t *testing.T) {
	program := `104,1125899906842624,99`
	memory := computer.LoadInstructionsWithMemoryAlloc(program, 8192)
	_, out := computer.Run(memory, []int64{0})
	fmt.Printf("out=%v\n", out)
	assert.Equal(t, int64(1125899906842624), out[0])
}

func TestMain(t *testing.T) {
	boostCode, err := ioutil.ReadFile("./boost_code_d9.txt")
	if err != nil {
		panic(err)
	}

	memory := computer.LoadInstructionsWithMemoryAlloc(string(boostCode), 8192)
	_, out := computer.Run(memory, []int64{1})
	fmt.Printf("out: %v", out)
	assert.Equal(t, int64(3507134798), out[0])
}
