package main

import (
	"os"

	"github.com/flyhard/nordea2ynab/convert"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		convert.Convert(arg)
	}

}
