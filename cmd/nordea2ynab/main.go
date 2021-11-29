package main

import (
	"nordea2Ynab/convert"
	"os"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		convert.Convert(arg)
	}

}
