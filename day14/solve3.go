package main

import (
	"fmt"
	"math"
)

func solve3(amount int, recipes map[string]reactionType) (int, map[string]int, int) {
	var oreNeeded = 0
	var supply = make(map[string]int)

	var orders = []elementType{
		{
			symbol: "FUEL",
			q:      amount,
		},
	}

	for len(orders) > 0 {
		order := orders[0]
		orders = orders[1:]

		if _, found := supply[order.symbol]; !found {
			supply[order.symbol] = 0
		}

		if order.symbol == "ORE" {
			oreNeeded += order.q
		} else if order.q <= supply[order.symbol] {
			supply[order.symbol] -= order.q
		} else {
			amountNeeded := order.q - supply[order.symbol]
			recipe := recipes[order.symbol]

			batches := int(math.Ceil(float64(amountNeeded) / float64(recipe.output.q)))
			for _, ingredient := range recipe.inputs {
				orders = append(orders, elementType{symbol: ingredient.symbol, q: ingredient.q * batches})
			}

			leftoverAmount := batches*recipe.output.q - amountNeeded
			supply[order.symbol] = leftoverAmount
		}

	}

	return oreNeeded, supply, 0
}

func solve3Part2(recipes map[string]reactionType) int {

	guess := -1
	upperBound := -1
	lowerBound := 1
	oreCapacity := 1000000000000

	for lowerBound+1 != upperBound {
		if upperBound == -1 {
			guess = lowerBound * 2
		} else {
			guess = (upperBound + lowerBound) / 2
		}

		oreNeeded, _, _ := solve3(guess, recipes)
		if oreNeeded > oreCapacity {
			upperBound = guess
		} else {
			lowerBound = guess
		}

		fmt.Println("oreNeeded", oreNeeded, guess)
	}

	return guess

}
