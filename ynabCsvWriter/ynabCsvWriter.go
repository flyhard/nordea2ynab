package ynabCsvWriter

import (
	"github.com/flyhard/nordea2ynab/csvTypes"
)

type YnabRecord struct {
	Date    csvTypes.DateTime `csv:"Date"`
	Payee   string            `csv:"Payee"`
	Memo    string            `csv:"Memo"`
	Outflow string            `csv:"Outflow"`
	Inflow  string            `csv:"Inflow"`
}
