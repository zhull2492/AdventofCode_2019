package day_01

import
(
	"fmt"
	"math"
	"../fileParse"
	"../zutils"
)

func calcFuel(mass int) (int) {
	return int(math.Floor(float64(mass) / float64(3))) - 2
}

func recursiveFuel(fuel int) (int) {
	if fuel > 0 {
		return fuel + recursiveFuel(calcFuel(fuel))
	} else {
		return 0
	}
}

func Run(fn string, debug bool) ([]int) {
	fmt.Println("Day 01")
	masses := zutils.String2Int(fileParse.ReadData(fn))

	totalFuel_A := 0
	totalFuel_B := 0

	for _, mass := range masses {
		totalFuel_A += calcFuel(mass)
		totalFuel_B += recursiveFuel(calcFuel(mass))
	}

	return []int{totalFuel_A, totalFuel_B}
}
