package convert

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flyhard/nordea2ynab/decimal"
)

const regex = "(?P<Memo>Autogiro|Överföring|Reservation Kortköp|Kortköp \\d+|Swish betalning|Swish återbetalning|Betalning BG \\d+-\\d+).(?P<Payee>.*)"

func Convert(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", filename, err)
	}

	r := csv.NewReader(file)
	r.Comma = ';'
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read file %q: %w", file.Name(), err)
	}
	for _, record := range records {
		fmt.Println(record)
	}
	w := new(strings.Builder)
	ext := filepath.Ext(filename)
	name := strings.TrimRight(filename, ext)
	outFileName := name + ".ynab.csv"
	if os.Truncate(outFileName, 0) != nil {
		log.Printf("failed to truncate: %v", err)
	}
	outFile, err := os.OpenFile(outFileName, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open out file %q: %w", outFileName, err)
	}
	out := csv.NewWriter(outFile)
	err = out.Write([]string{"Date", "Payee", "Memo", "Outflow", "Inflow"})
	if err != nil {
		return fmt.Errorf("cannot write: %w", err)
	}
	re, err := regexp.Compile(regex)
	if err != nil {
		return fmt.Errorf("failed to parse regex %q: %w", regex, err)
	}
	reserved, err := regexp.Compile("Reservation Kortköp")
	if err != nil {
		return fmt.Errorf("failed to parse regex: %w", err)
	}
	for i, record := range records {
		if i == 0 {
			continue
		}
		err := handleLoop(record, out, re, reserved)
		if err != nil {
			return err
		}
	}
	out.Flush()
	fmt.Println(w.String())
	return nil
}

func handleLoop(record []string, out *csv.Writer, re *regexp.Regexp, reserved *regexp.Regexp) error {
	date := strings.TrimSpace(record[0])
	payee := ""
	memo := strings.TrimSpace(record[5])

	if re.MatchString(memo) {
		submatch := re.FindStringSubmatch(memo)
		payee = submatch[re.SubexpIndex("Payee")]
		memo = submatch[re.SubexpIndex("Memo")]
	}
	if reserved.MatchString(memo) {
		return nil
	}
	amtText := strings.TrimSpace(record[1])
	money, err := decimal.NewFromString(amtText)
	if err != nil {
		return fmt.Errorf("failed to parse %q: %w", record, err)
	}
	outflow := "0"
	inflow := "0"
	if money.IsNegative() {
		outflow = fmt.Sprintf("%.2f", money.Absolute().AsMajorUnits())
	} else {
		inflow = fmt.Sprintf("%.2f", money.Absolute().AsMajorUnits())
	}
	outRec := []string{date, payee, memo, outflow, inflow}
	log.Println(outRec)
	if out.Write(outRec) != nil {
		return fmt.Errorf("failed to write: %w", err)
	}
	return nil
}
