package main

import
(
	"flag"
	"fmt"
	"os"
	"strconv"
	"./zutils"
	"./day_1"
	"./day_2"
	"./day_3"
	"./day_4"
	"./day_5"
	"./day_6"
	"./day_7"
	"./day_8"
	"./day_9"
	"./day_10"
	"./day_11"
	"./day_12"
	"./day_13"
	"./day_14"
	"./day_15"
	"./day_16"
	"./day_17"
	"./day_18"
	"./day_19"
	"./day_20"
	"./day_21"
	"./day_22"
	"./day_23"
	"./day_24"
	"./day_25"
)

func main() {

	dayRun := flag.Int("d", 0, "Day Number to Execute")
	parseFile := flag.String("f", "", "Input File to parse")
	validateResult := flag.Bool("c", false, "Check the results")
	debugFlag := flag.Bool("b", false, "Debug Flag")

	flag.Parse()

	if *dayRun == 0 {
		fmt.Println(
`Welcome to Advent of Code.
Please execute with valid day number.`)
		os.Exit(0)
	} else if *dayRun > 25 || *dayRun < 0 {
		fmt.Println("Invalid Day Number")
		os.Exit(1)
	}

	var iresult []int
	var sresult []string
	var fresult []float64

	switch *dayRun {
	case 1:
		iresult = day_01.Run(*parseFile, *debugFlag)
	case 2:
		day_02.Run()
	case 3:
		day_03.Run()
	case 4:
		day_4.Run()
	case 5:
		day_5.Run()
	case 6:
		day_6.Run()
	case 7:
		day_7.Run()
	case 8:
		day_8.Run()
	case 9:
		day_9.Run()
	case 10:
		day_10.Run()
	case 11:
		day_11.Run()
	case 12:
		day_12.Run()
	case 13:
		day_13.Run()
	case 14:
		day_14.Run()
	case 15:
		day_15.Run()
	case 16:
		day_16.Run()
	case 17:
		day_17.Run()
	case 18:
		day_18.Run()
	case 19:
		day_19.Run()
	case 20:
		day_20.Run()
	case 21:
		day_21.Run()
	case 22:
		day_22.Run()
	case 23:
		day_23.Run()
	case 24:
		day_24.Run()
	case 25:
		day_25.Run()
	}


	zutils.PrintResults(iresult, fresult, sresult)

	if *validateResult {
		chkfile := "./day_" + strconv.Itoa(*dayRun) + "/accepted.txt"
		zutils.ValidateResults(chkfile, iresult, fresult, sresult)
	}
}
