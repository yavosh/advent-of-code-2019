package main

import "fmt"

func copymap(in map[string]int) map[string]int {

	res := make(map[string]int, 0)
	for key, val := range in {
		res[key] = val
	}

	return res
}

func keys(in map[string]int) []string {
	res := make([]string, 0)
	for key := range in {
		res = append(res, key)
	}
	return res
}

func solve2(reactions map[string]reactionType) (int, map[string]int, map[string]int) {
	var work = make(map[string]int)
	var backpack = make(map[string]int)
	var excess = make(map[string]int)
	var padded = make(map[string]int)

	var done = false

	// ore := 0

	for _, input := range reactions["FUEL"].inputs {
		work[input.symbol] = input.q
	}

	loops := 0

	for !done {

		fmt.Printf("Loop %d\n", loops)

		loops++

		in := keys(work)

		oreCount := 0
		notOreCount := 0

		oreReactionCount := 0
		notOreReactionCount := 0

		for _, key := range in {
			fmt.Printf("key %s len %d\n", key, len(in))

			if key == "ORE" {
				oreCount++
			} else {
				notOreCount++
			}

			if len(in) == 1 {
				// only ore
				done = true
			}

			reaction := reactions[key]
			if reaction.output.symbol == "" {
				// this is ore
				continue
			}

			if reaction.OreReaction() {
				oreReactionCount++
			} else {
				fmt.Printf("notOreReactionCount %v \n", reaction)
				notOreReactionCount++
			}

			if work[key]%reaction.output.q == 0 {
				multiplier := work[key] / reaction.output.q
				for _, in := range reaction.inputs {
					found := work[in.symbol]
					work[in.symbol] = found + (in.q * multiplier)
				}

				delete(work, key)
				continue
			}

			if work[key] > reaction.output.q {
				multiplier := work[key] / reaction.output.q
				remainder := work[key] % reaction.output.q
				fmt.Printf("el=%s m=%d r=%d\n", key, multiplier, remainder)

				for _, in := range reaction.inputs {
					found := work[in.symbol]
					work[in.symbol] = found + (in.q * multiplier)
				}

				backpackFound := backpack[key]
				backpack[key] = backpackFound + remainder

				delete(work, key)
				continue
			}

			if work[key] < reaction.output.q {
				fmt.Printf("**** Less %s\n", key)
				//backpack[key] = work[key]
				//delete(work, key)
			}

		}

		fmt.Printf("oreCount=%d notOreCount=%d\n", oreCount, notOreCount)
		fmt.Printf("oreReactionCount=%d notOreReactionCount=%d\n", oreReactionCount, notOreReactionCount)
		fmt.Printf("in=%v\n", in)

		if oreReactionCount == 0 && notOreReactionCount > 0 {
			// Run out of work but not done
			fmt.Printf("************ PAD reactions: \n")

			for key, val := range work {
				if key == "ORE" {
					continue
				}

				fmt.Printf("************ PAD key %s %d\n", key, val)
				reaction := reactions[key]

				if reaction.OreReaction() {
					// do not pad reactions that resolve to ore
					continue
				}

				work[key] = reaction.output.q
				paddedFound := padded[key]
				padded[key] = paddedFound + (reaction.output.q - val)
			}
		}

		done = done || loops > 1000

		// 322
	}

	fmt.Printf("work %+v\n", work)
	fmt.Printf("excess %+v\n", excess)
	fmt.Printf("backpack %+v\n", backpack)
	fmt.Printf("padded %+v\n", padded)

	for key, val := range backpack {

		reaction := reactions[key]

		if len(reaction.inputs) != 1 && reaction.inputs[0].symbol != "ORE" {
			panic(fmt.Sprintf("Expecting ore got %+v", reaction))
		}

		ore := reaction.inputs[0].q
		found := backpack[key]

		remainder := ore - val

		work["ORE"] += ore
		foundExcess := excess[key]
		excess[key] = foundExcess + remainder

		fmt.Printf("found: %d %v %v\n", found, key, val)
		fmt.Printf("found: %v\n", reaction)
	}

	fmt.Printf("work %+v\n", work)
	fmt.Printf("excess %+v\n", excess)
	fmt.Printf("backpack %+v\n", backpack)
	fmt.Printf("padded %+v\n", padded)

	return work["ORE"], work, excess
}
