package zutils

import (
	"../fileParse"
	"strconv"
	"fmt"
)

func checkInts (calcVal []int, accepted string) {

	accepted_Vals := String2Int(fileParse.ReadData(accepted))

	if calcVal[0] == accepted_Vals[0] && calcVal[1] == accepted_Vals[1] {
                fmt.Println("The results have been verified")
        } else {
                fmt.Println("Error -- Mismatched Results")
        }
}

func String2Int (invals []string) (outvals []int) {

	for _, ent := range invals {
		tmpval, _ := strconv.Atoi(ent)
		outvals = append(outvals, tmpval)
	}

	return outvals
}

func PrintResults (ires []int, fres []float64, sres []string) {

	if len(ires) > 0 {
		fmt.Printf("Part A: %d\nPart B: %d\n", ires[0], ires[1])
	} else if len(fres) > 0 {
		fmt.Printf("Part A: %f\nPart B: %f\n", fres[0], fres[1])
	} else if len(sres) > 0 {
		fmt.Printf("Part A: %s\nPart B: %s\n", sres[0], sres[1])
	}
}

func ValidateResults (chkfile string, ires []int, fres []float64, sres []string) {

	if len(ires) > 0 {
		checkInts(ires, chkfile)
	} else if len(fres) > 0 {
		fmt.Println("Float Results")
	} else if len(sres) > 0 {
		fmt.Println("String Results")
	}
}
