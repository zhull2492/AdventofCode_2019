package fileParse

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func split(fn string, sep string) (splitdat []string) {
	for _, sl := range strings.Split(fn, sep) {
		if sl != "" {
			splitdat = append(splitdat, sl)
		}
	}
	return
}

func parse(fn string) (strdat string) {
	dat, err := ioutil.ReadFile(fn)

	if err != nil {
		fmt.Printf("Read File Error on %s\n", fn)
		os.Exit(2)
	}

	strdat = string(dat)
	return strdat
}

func ReadData(fn string) ([]string) {
	return split(parse(fn), "\n")
}
