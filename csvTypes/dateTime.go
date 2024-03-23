package csvTypes

import (
	"fmt"
	"time"
)

type DateTime struct {
	time.Time
	IsReserved bool
}

const reserved = "Reserverat"

// MarshalCSV Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	if date.IsReserved {
		return reserved, nil
	}
	return date.Time.Format("2006/01/02"), nil
}

// You could also use the standard Stringer interface
func (date *DateTime) String() string {
	return date.String() // Redundant, just for example
}

func (date *DateTime) UnmarshalCSV(value string) error {
	if value == reserved {
		date.IsReserved = true
		return nil
	}
	parse, err := time.Parse("2006/01/02", value)
	if err != nil {
		return fmt.Errorf("failed to parse date: %v", err)
	}
	date.Time = parse
	return nil
}
