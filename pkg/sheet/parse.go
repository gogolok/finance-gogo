package sheet

import (
	"strconv"
	"time"
)

func parseAll(d []SourceRawData) ([]Sheet, error) {
	sheets := []Sheet{}

	for _, entry := range d {
		sheet, err := parse(entry)
		if err != nil {
			return sheets, err
		}
		sheets = append(sheets, *sheet)
	}

	return sheets, nil
}

func parse(d SourceRawData) (*Sheet, error) {
	sheet := &Sheet{
		Name:    d.Name,
		Entries: []Entry{},
	}

	records := d.Records[1:] // get rid of the header line
	for _, record := range records {
		// skip every `empty` line
		if record[0] == "" {
			continue
		}

		// Date
		const shortForm = "02.01.2006"
		date, err := time.Parse(shortForm, record[1])
		if err != nil {
			return nil, err
		}

		// Paid
		paid, err := strconv.ParseFloat(record[2], 10)
		if err != nil {
			return nil, err
		}

		// Costs
		if record[3] == "" {
			record[3] = "0.00"
		}
		provision, err := strconv.ParseFloat(record[3], 10)
		if err != nil {
			return nil, err
		}
		if record[4] == "" {
			record[4] = "0.00"
		}
		clearstream, err := strconv.ParseFloat(record[4], 10)
		if err != nil {
			return nil, err
		}
		if record[5] == "" {
			record[5] = "0.00"
		}
		market, err := strconv.ParseFloat(record[5], 10)
		if err != nil {
			return nil, err
		}
		costs := provision + clearstream + market

		// Amount
		amount, err := strconv.ParseFloat(record[7], 10)
		if err != nil {
			return nil, err
		}

		entry := Entry{
			Source: record[0],
			WKN:    record[6],
			Date:   date,
			Paid:   paid,
			Costs:  costs,
			Amount: amount,
		}
		sheet.Entries = append(sheet.Entries, entry)
	}

	return sheet, nil
}
