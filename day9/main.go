package main

import (
	"fmt"
	"io/ioutil"

	"github.com/yavosh/advent-of-code-2019/computer"
)

func main() {
	boostCode, err := ioutil.ReadFile("./boost_code_d9.txt")
	if err != nil {
		panic(err)
	}

	memory := computer.LoadInstructionsWithMemoryAlloc(string(boostCode), 8192)
	_, out := computer.Run(memory, []int64{2})
	fmt.Printf("out: %v", out)
}
