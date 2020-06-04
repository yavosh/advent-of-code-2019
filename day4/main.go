package main

import (
	"fmt"
	"strconv"
)

var (
	input = []int{234208, 765869}
)

func validPass(pass string) (bool, string) {
	//fmt.Printf("validPass: %s\n", pass)

	if len(pass) != 6 {
		return false, "length"
	}

	if _, err := strconv.Atoi(pass); err != nil {
		// not a number
		return false, "format"
	}

	doublesCount := 0
	consecutive := 0

	for pos := 1; pos < len(pass); pos++ {
		prev, _ := strconv.Atoi(string(pass[pos-1]))
		curr, _ := strconv.Atoi(string(pass[pos]))

		if prev == curr {
			if consecutive == 0 {
				consecutive = 2
			} else {
				consecutive++
			}
		} else {
			if consecutive == 2 {
				doublesCount++
			}

			consecutive = 0
		}

		if curr < prev {
			return false, "decreasing"
		}

		// fmt.Printf("Letter: %c %c %d->%d\n", pass[pos-1], pass[pos], prev, curr)
	}

	if consecutive == 2 {
		doublesCount++
	}

	if doublesCount == 0 {
		return false, "no-double"
	}

	return true, "ok"
}

func main() {

	validPasswords := make([]string, 0)
	for pass := input[0]; pass <= input[1]; pass++ {
		passValue := fmt.Sprintf("%d", pass)
		valid, _ := validPass(passValue)
		if valid {
			validPasswords = append(validPasswords, passValue)
		}
	}

	fmt.Printf("valid passwords len=%d pass=%v", len(validPasswords), validPasswords)
}
