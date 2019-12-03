package day_02

import
(
	"fmt"
	"os"
	"../fileParse"
	"../zutils"
)

func moveIdx(idx int, numele int) (newidx int) {

	newidx = idx + 1
	if newidx == numele {
		newidx = 0
	}
	return newidx

}

func runCode (codes []int) (int) {
	var opcode int
	var reg1 int
	var reg2 int
	var dest int
	idx := 0
	cloop := true

	for cloop == true {
		opcode = codes[idx]
		idx = moveIdx(idx, len(codes))
		reg1 = codes[idx]
		idx = moveIdx(idx, len(codes))
		reg2 = codes[idx]
		idx = moveIdx(idx, len(codes))
		dest = codes[idx]
		idx = moveIdx(idx, len(codes))
		switch opcode {
		case 1:
			codes[dest] = codes[reg1] + codes[reg2]
		case 2:
			codes[dest] = codes[reg1] * codes[reg2]
		case 99:
			cloop = false
		default:
			fmt.Println("Runtime Error")
			os.Exit(2)
		}
	}
	return codes[0]
}

func Run(fn string) ([]int) {
	fmt.Println("Day 02")

	vals := fileParse.ReadData(fn)
	codes := zutils.String2Int(zutils.String2List(vals[0], ","))

	var origcodes = make([]int, len(codes))
	copy(origcodes, codes)

	codes[1] = 12
	codes[2] = 2

	final_A := runCode(codes)

	found := false
	target := 19690720

	var noun int
	var verb int

	for noun = 0;  noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			copy(codes, origcodes)

			codes[1] = noun
			codes[2] = verb

			if runCode(codes) == target {
				found = true
				break
			}
		}
		if found == true {
			break
		}
	}

	return []int{final_A, (100*noun)+verb}

}
