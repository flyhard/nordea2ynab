package nordeaCsvReader

import (
	"encoding/json"
	"testing"
)

func TestReadingData(t *testing.T) {
	filename := "testdata/anonymized_transactions.csv"
	transactions, err := ReadNordeaTransactionsFromFile(filename)
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}
	t.Logf("Read %v transactions", len(*transactions))
	if len(*transactions) != 10 {
		t.Errorf("Failed to convert right amount of rows")
		for _, transaction := range *transactions {
			marshal, _ := json.Marshal(transaction)
			t.Logf("%v", string(marshal))
		}
	}
}
