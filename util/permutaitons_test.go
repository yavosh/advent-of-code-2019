package util_test

import (
	"fmt"
	"testing"

	"github.com/yavosh/advent-of-code-2019/util"
)

func TestMakePermutations(t *testing.T) {
	fmt.Printf("[][]int{ \n")
	for _, permutations := range util.Permutations([]int{5, 6, 7, 8, 9}) {
		fmt.Printf("{")
		for _, value := range permutations {
			fmt.Printf("%d,", value)
		}
		fmt.Printf("},\n")
	}
	fmt.Printf("}\n")
	t.Fail()
}
