package nordeaCsvReader

import (
	"encoding/csv"
	"fmt"
	"github.com/flyhard/nordea2ynab/csvTypes"
	"github.com/gocarina/gocsv"
	"io"
	"os"
)

// Bokföringsdag;Belopp;Avsändare;Mottagare;Namn;Rubrik;Saldo;Valuta;

type NordeaTransaction struct {
	BookingDate csvTypes.DateTime `csv:"Bokföringsdag"`
	Amount      string            `csv:"Belopp"`
	FromAccount string            `csv:"Avsändare"`
	To          string            `csv:"Mottagare"`
	Name        string            `csv:"Namn"`
	Category    string            `csv:"Rubrik"`
	Saldo       string            `csv:"Saldo"`
	Currency    string            `csv:"Valuta"`
	Empty       string            `csv:""`
}

func ReadNordeaTransactionsFromFile(filename string) (*[]NordeaTransaction, error) {
	file, err := os.Open(filename)
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use semicolon as delimiter
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q: %w", filename, err)
	}

	transactions := &[]NordeaTransaction{}
	if err := gocsv.Unmarshal(file, transactions); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal: %w", err)
	}

	return transactions, nil
}
