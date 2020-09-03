package main

import "fmt"

func solve(reactions map[string]reactionType) (int, map[string]int, map[string]int) {
	var backpack = make(map[string]int)
	var excess = make(map[string]int)
	var work = make([]elementType, 0)

	ore := 0
	begin := reactions["FUEL"]
	work = append(work, begin.output)

	for len(work) > 0 {
		fmt.Printf("work=%+v backpack=%+v ore=%d \n", work, backpack, ore)

		element := work[0]
		work = work[1:]

		if len(work) == 0 {
			// finished data check if we can use backpack
			for element, amount := range backpack {
				elementReaction := reactions[element]

				if amount <= 0 {
					continue
				}

				if elementReaction.inputs[0].symbol == "ORE" {
					continue
				}

				fmt.Printf("&BK element=%s amount=%d elementReaction=%v\n", element, amount, elementReaction)

				multiplier := amount / elementReaction.output.q
				remainder := amount % elementReaction.output.q

				if remainder > 0 {
					// add to backpack
					multiplier = multiplier + 1
				}

				for _, input := range elementReaction.inputs {
					workInput := elementType{symbol: input.symbol, q: input.q * multiplier}
					work = append(work, workInput)
				}

				backpack[element] = backpack[element] - amount
				excess[element] = multiplier*elementReaction.output.q - amount
			}
		}

		if element.symbol == "ORE" {
			ore += element.q
			continue
		}

		if _, found := backpack[element.symbol]; !found {
			backpack[element.symbol] = 0
		}

		elementReaction := reactions[element.symbol]
		if elementReaction.output.q > element.q {
			backpack[element.symbol] += element.q
			fmt.Printf("add backpack %s %d backpack=%+v\n", element.symbol, element.q, backpack)

			//continue
			//if elementReaction.output.q <= backpack[element.symbol] {
			//	backpack[element.symbol] -= elementReaction.output.q
			//	fmt.Printf("remove backpack %s %d backpack=%+v\n",
			//		elementReaction.output.symbol, elementReaction.output.q, backpack)
			//
			//	workInput := elementType{symbol: input.symbol, q: input.q}
			//	work = append(work, workInput)
			//	fmt.Printf("@@@ workInput=%+v\n", workInput)
			//}

			continue
		}

		multiplier := element.q / elementReaction.output.q
		remainder := element.q % elementReaction.output.q
		fmt.Printf("^^ multiplier=%d remainder=%d\n", multiplier, remainder)

		if remainder > 0 {
			// add to backpack
			backpack[element.symbol] += remainder
		}

		for _, input := range elementReaction.inputs {
			fmt.Printf("^ input %+v element %+v elementReaction %+v\n", input, element, elementReaction)
			workInput := elementType{symbol: input.symbol, q: input.q * multiplier}
			work = append(work, workInput)
		}
	}

	// 180697
	// 181392
	// 695 695/139 = 5 , 4NVRVD
	// 2220043
	// 2210736
	// 9307

	fmt.Printf("PRE work=%+v backpack=%+v excess=%+v \n", work, backpack, excess)
	for element, amount := range backpack {

		if amount == 0 {
			continue
		}

		// purge backpack
		backpack[element] -= amount
		elementReaction := reactions[element]
		fmt.Printf("&element=%s amount=%d elementReaction=%v ore=%d\n", element, amount, elementReaction, ore)
		for _, input := range elementReaction.inputs {

			if input.symbol != "ORE" {
				panic("Expecting ore")
			}

			multiplier := amount / elementReaction.output.q
			remainder := amount % elementReaction.output.q
			if remainder > 0 {
				multiplier = multiplier + 1
				remainder = elementReaction.output.q - remainder
			}

			ore += input.q * multiplier
			if _, found := excess[element]; found {
				excess[element] += remainder
			} else {
				excess[element] = remainder
			}
		}
	}

	fmt.Printf("DONE work=%+v backpack=%+v excess=%+v \n", work, backpack, excess)
	return ore, backpack, excess
}
