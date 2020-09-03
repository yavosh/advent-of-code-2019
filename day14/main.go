package main

import (
	"fmt"
	"strconv"
	"strings"
)

type elementType struct {
	symbol string
	q      int
}

type reactionType struct {
	output elementType
	inputs []elementType
}

func (t *reactionType) OreReaction() bool {
	fmt.Printf("T %+v\n", t)
	return len(t.inputs) > 0 && t.inputs[0].symbol == "ORE"
}

type treeNode struct {
	element  string
	count    int
	branches []treeNode
}

func NewElementFromString(in string) elementType {
	parts := strings.Split(in, " ")

	var symbol = strings.TrimSpace(parts[1])
	quantity, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	return elementType{
		symbol: symbol,
		q:      quantity,
	}
}

func loadInput(inputString string) map[string]reactionType {
	lines := strings.Split(inputString, "\n")
	var reactions = make(map[string]reactionType, 0)
	for _, line := range lines {
		args := strings.Split(line, "=>")
		inputs := strings.Split(args[0], ",")
		output := strings.TrimSpace(args[1])
		reactionInputs := make([]elementType, 0)

		for _, input := range inputs {
			reactionInputs = append(reactionInputs, NewElementFromString(strings.TrimSpace(input)))
		}

		reaction := reactionType{
			output: NewElementFromString(output),
			inputs: reactionInputs,
		}

		reactions[reaction.output.symbol] = reaction
	}

	return reactions
}

func main() {
	var _ = `10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`

	var _ = `9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL`

	var _ = `157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`

	var _ = `2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF`

	var _ = `171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`

	var input = `1 FJFL, 1 BPVQN => 7 CMNH
6 FJFL, 2 KZJLT, 3 DZQJ => 2 NSPZ
11 TPZDN => 2 TNMC
1 NSPZ, 2 KQVL => 2 HPNWP
3 XHDVT => 9 LRBN
3 LRBN => 6 TPZDN
1 KPFLZ, 1 XVXCZ => 6 WHMLV
1 BDWQP, 1 JPGK, 1 MTWG => 5 GLHWQ
2 BGLTP, 1 HPNWP, 2 GLHWQ, 9 CRJZ, 22 QVQJ, 3 PHGWC, 1 BDWQP => 3 LKPNB
4 BDSB => 2 PNSD
2 BRJDJ, 13 THQR => 2 BGLTP
1 WHJKH, 2 JBTJ => 6 THQR
1 JBTJ => 9 WGVP
10 CTXHZ, 2 DGMN => 5 TNVC
7 LCSV, 1 LKPNB, 36 CMNH, 1 JZXPH, 20 DGJPN, 3 WDWB, 69 DXJKC, 3 WHJKH, 18 XSGP, 22 CGZL, 2 BNVB, 57 PNSD => 1 FUEL
13 CRCG, 8 NMQN => 1 JZXPH
2 FZVS, 2 ZPFBH => 9 SRPD
1 QPNTQ, 4 QVQJ, 1 XZKTG => 9 WDWB
6 SXZW => 5 FJFL
6 GVGZ => 6 ZPFBH
1 JPGK, 3 WDFXH, 22 FJFL => 7 BDSB
3 WHMLV => 4 JPGK
7 CGZL, 4 LRBN => 8 MTWG
11 SXZW, 33 ZTBFN => 4 XVXCZ
1 FZVS, 1 TNMC, 7 JPGK => 9 FLHW
2 XKFZ => 8 CGZL
5 WHMLV => 8 MQRS
1 QVSH, 6 TPZDN, 9 JQHCH => 2 BMNJ
3 CMNH, 10 XKFZ => 2 KQVL
119 ORE => 9 PSPQ
1 WGVP, 18 BRJDJ => 9 PHGWC
110 ORE => 6 NMQN
13 NMQN, 24 XVXCZ, 9 XHDVT => 6 KQVS
6 TNMC => 4 DXJKC
1 XZKTG => 8 WHJKH
1 KPFLZ, 1 LRBN, 7 KQVS => 9 JBTJ
1 XKFZ => 8 JVGD
152 ORE => 7 SXZW
1 BDWQP => 5 CTXHZ
2 JVGD, 8 DGMN, 2 MTWG => 6 QVQJ
1 KQVL => 2 BNVB
3 DZQJ, 37 MQRS => 4 CRJZ
1 KQVL, 5 WDFXH => 7 BDWQP
3 GVGZ => 3 BPVQN
4 PSPQ, 6 ZTBFN => 1 KPFLZ
34 FBTG => 9 XZKTG
14 TNMC, 4 FZVS, 3 MTWG => 9 KZJLT
157 ORE => 6 GVGZ
5 JVGD, 11 JPGK => 5 CRCG
1 SXZW, 1 NMQN => 3 XHDVT
1 FBTG => 5 JQHCH
3 WDFXH, 4 FZVS, 9 CGFML => 6 DZQJ
5 BDWQP, 3 TNVC, 7 SRPD, 1 WDFXH, 3 JQHCH, 4 QVQJ, 5 CRCG, 4 DGMN => 7 XSGP
1 KPFLZ, 3 TPZDN, 1 SRPD => 6 FBTG
1 WHMLV, 3 BDSB, 2 JVGD => 9 LCSV
13 XZKTG => 4 QVSH
1 XHDVT => 7 XKFZ
1 CMNH, 1 KQVS, 2 XVXCZ => 6 CGFML
6 FLHW => 4 BRJDJ
2 KQVL, 2 WGVP, 7 BMNJ, 11 KQVS, 1 HPNWP, 6 CRJZ => 4 DGJPN
2 DZQJ, 1 BDSB => 2 DGMN
1 XVXCZ, 4 MQRS => 3 WDFXH
5 FLHW, 10 JPGK, 1 XZKTG => 4 QPNTQ
2 LRBN => 9 FZVS
149 ORE => 8 ZTBFN`

	reactions := loadInput(input)

	ore, backpack, excess := solve3(1, reactions)
	fmt.Printf("ore=%d backpack=%+v excess=%+v \n", ore, backpack, excess)

	guess := solve3Part2(reactions)
	fmt.Println("guess:", guess)

	//	treeRoot := buildCollapsedTree(nodes["FUEL"], nodes, 1)
	//	printTree(treeRoot, 1)
	// ore=158482
}
