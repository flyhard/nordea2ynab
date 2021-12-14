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
	"github.com/flyhard/nordea2ynab/ofx"
)

const regex = "(?P<Memo>Autogiro|Överföring|Reservation Kortköp|Kortköp \\d+|Swish betalning|Swish återbetalning|Betalning BG \\d+-\\d+).(?P<Payee>.*)"

func Qif(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %q: %v", filename, err)
	}

	r := csv.NewReader(file)
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		fmt.Println(record)
	}
	w := new(strings.Builder)
	ext := filepath.Ext(filename)
	name := strings.TrimRight(filename, ext)
	qifFile := name + ".qif"
	if os.Truncate(qifFile, 0) != nil {
		log.Printf("failed to truncate: %v", err)
	}
	out, err := os.Create(qifFile)
	if err != nil {
		return err
	}
	defer out.Close()
	statementTransactions := []*ofx.StatementTransaction{}
	bankTransactionList := &ofx.BankTransactionList{
		STMTTRN: statementTransactions,
	}
	pendingTransactions := []*ofx.PendingTransaction{}
	pendingTansactionList := &ofx.PendingTransactionList{
		STMTTRNP: pendingTransactions,
	}
	statementResponse := &ofx.StatementResponse{
		CURDEF:        "SEK",
		BANKTRANLIST:  bankTransactionList,
		BANKTRANLISTP: pendingTansactionList,
	}
	stmnt := &ofx.StatementTransactionResponse{
		STMTRS: statementResponse,
	}
	var statements []*ofx.StatementTransactionResponse
	bankResponseMessageSetV1 := &ofx.BankResponseMessageSetV1{
		STMTTRNRS: statements,
	}
	ofxData := ofx.OFX{
		BANKMSGSRSV1: bankResponseMessageSetV1,
	}

	////////
	reserved, err := regexp.Compile("Reservation.*")
	if err != nil {
		log.Fatalf("Failed to parse regex: %v", err)
	}
	re, err := regexp.Compile(regex)
	if err != nil {
		log.Fatalf("Failed to parse regex %q: %v", regex, err)
	}
	for i, record := range records {
		if i == 0 {
			continue
		}

		date := strings.TrimSpace(record[0])
		payee := ""
		memo := strings.TrimSpace(record[5])

		if re.MatchString(memo) {
			submatch := re.FindStringSubmatch(memo)
			payee = submatch[re.SubexpIndex("Payee")]
			memo = submatch[re.SubexpIndex("Memo")]
		}
		amtText := strings.TrimSpace(record[1])
		money, err := decimal.NewFromString(amtText)
		if err != nil {
			log.Fatalf("failed to parse %q: %v", record, err)
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
		if reserved.MatchString(memo) {
			ptrn := ofx.PendingTransaction{
				TRNTYPE:      "DEBIT",
				DTTRAN:       "",
				DTEXPIRE:     "",
				TRNAMT:       "",
				REFNUM:       "",
				NAME:         "",
				EXTDNAME:     "",
				MEMO:         "",
				IMAGEDATA:    nil,
				CURRENCY:     nil,
				ORIGCURRENCY: nil,
			}
		} else {
		}
	}
	out.Flush()
	fmt.Println(w.String())
}
func ConvertCsv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %q: %v", filename, err)
	}

	r := csv.NewReader(file)
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
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
		log.Fatalf("Failed to open out file %q: %v", outFileName, err)
	}
	out := csv.NewWriter(outFile)
	err = out.Write([]string{"Date", "Payee", "Memo", "Outflow", "Inflow"})
	if err != nil {
		log.Fatalf("cannot write: %v", err)
	}
	re, err := regexp.Compile(regex)
	if err != nil {
		log.Fatalf("Failed to parse regex %q: %v", regex, err)
	}
	for i, record := range records {
		if i == 0 {
			continue
		}
		date := strings.TrimSpace(record[0])
		payee := ""
		memo := strings.TrimSpace(record[5])

		if re.MatchString(memo) {
			submatch := re.FindStringSubmatch(memo)
			payee = submatch[re.SubexpIndex("Payee")]
			memo = submatch[re.SubexpIndex("Memo")]
		}
		amtText := strings.TrimSpace(record[1])
		money, err := decimal.NewFromString(amtText)
		if err != nil {
			log.Fatalf("failed to parse %q: %v", record, err)
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
			log.Fatalf("failed to write: %v", err)
		}
	}
	out.Flush()
	fmt.Println(w.String())
}
