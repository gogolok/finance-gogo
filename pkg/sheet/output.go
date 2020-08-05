package sheet

import (
	"encoding/json"
	"os"
)

func outputSheets(sheets []Sheet) error {
	b, err := json.MarshalIndent(sheets, "", "  ")
	if err != nil {
		return err
	}

	os.Stdout.Write(b)

	return nil
}

func outputGroupEntries(groupEntries []GroupEntry) error {
	b, err := json.MarshalIndent(groupEntries, "", "  ")
	if err != nil {
		return err
	}

	os.Stdout.Write(b)

	return nil
}
