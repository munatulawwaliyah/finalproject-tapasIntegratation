package model

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func CsvToSlice(data string) (map[string][]string, error) {
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return nil, fmt.Errorf("no data found in CSV")
	}

	table := make(map[string][]string)
	headers := records[0]

	for _, header := range headers {
		table[header] = []string{}
	}

	for _, record := range records[1:] {
		for i, value := range record {
			table[headers[i]] = append(table[headers[i]], value)
		}
	}
	return table, nil
}
