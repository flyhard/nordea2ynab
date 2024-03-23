package convert

import (
	"encoding/json"
	"fmt"
	"github.com/flyhard/nordea2ynab/csvTypes"
	"github.com/flyhard/nordea2ynab/nordeaCsvReader"
	"github.com/flyhard/nordea2ynab/ynabCsvWriter"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/flyhard/nordea2ynab/decimal"
)

const regex = "(?P<Memo>Autogiro|Överföring|Reservation Kortköp|Kortköp \\d+|Swish betalning|Swish återbetalning|Betalning BG \\d+-\\d+).(?P<Payee>.*)"

func Convert(filename string) error {
	nordeaTransactions, err := nordeaCsvReader.ReadNordeaTransactionsFromFile(filename)
	if err != nil {
		return err
	}
	re, err := regexp.Compile(regex)
	if err != nil {
		return fmt.Errorf("failed to parse regex %q: %w", regex, err)
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

	ynabRecords := &[]ynabCsvWriter.YnabRecord{}
	for _, record := range *nordeaTransactions {
		if record.BookingDate.IsReserved {
			continue
		}
		ynabRecord, err := handleLoop(record, re)
		if err != nil {
			return fmt.Errorf("failed to handle row: %w", err)
		}
		*ynabRecords = append(*ynabRecords, *ynabRecord)
	}
	err = gocsv.Marshal(ynabRecords, outFile)
	if err != nil {
		return fmt.Errorf("failed to write csv file %v: %w", outFile.Name(), err)
	}
	fmt.Println(w.String())
	return nil
}

func handleLoop(record nordeaCsvReader.NordeaTransaction, re *regexp.Regexp) (*ynabCsvWriter.YnabRecord, error) {
	payee := ""
	memo := strings.TrimSpace(record.Category)

	if re.MatchString(memo) {
		submatch := re.FindStringSubmatch(memo)
		payee = submatch[re.SubexpIndex("Payee")]
		memo = submatch[re.SubexpIndex("Memo")]
	}
	amtText := strings.TrimSpace(record.Amount)
	money, err := decimal.NewFromString(amtText)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %v: %w", record, err)
	}
	outflow := "0"
	inflow := "0"
	if money.IsNegative() {
		outflow = fmt.Sprintf("%.2f", money.Absolute().AsMajorUnits())
	} else {
		inflow = fmt.Sprintf("%.2f", money.Absolute().AsMajorUnits())
	}
	ynabRecord := ynabCsvWriter.YnabRecord{
		Date:    csvTypes.DateTime{Time: record.BookingDate.Time},
		Payee:   payee,
		Memo:    memo,
		Outflow: outflow,
		Inflow:  inflow,
	}
	marshal, _ := json.Marshal(ynabRecord)
	fmt.Printf("%v\n", string(marshal))

	return &ynabRecord, nil
}
