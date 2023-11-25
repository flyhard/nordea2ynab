package main

import (
	"log"
	"os"
	"strings"

	"github.com/flyhard/nordea2ynab/convert"
)

func main() {
	for i, filename := range os.Args {
		if i == 0 {
			continue
		}
		if strings.Contains(filename, ".ynab.") {
			log.Printf("Skipping converted file %q", filename)
			continue
		}
		log.Printf("Processing file %s", filename)
		err := convert.Convert(filename)
		if err != nil {
			log.Fatalf("Failed to convert: %v", err)
		}
		err = os.Remove(filename)
		if err != nil {
			log.Fatalf("Failed to remove: %v", err)
			return
		}
	}

}
