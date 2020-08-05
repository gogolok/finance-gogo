package sheet

func Data() ([]Sheet, error) {
	config, err := ParseConfig()
	if err != nil {
		return []Sheet{}, err
	}

	sourceRawData, err := fetchAllSourceRawData(config.Sources)
	if err != nil {
		return []Sheet{}, err
	}

	sheets, err := parseAll(sourceRawData)
	if err != nil {
		return []Sheet{}, err
	}

	return sheets, nil
}

func FetchDataAndOutput() error {
	sheets, err := Data()
	if err != nil {
		return err
	}

	return outputSheets(sheets)
}

func GroupDataAndOutput() error {
	sheets, err := Data()
	if err != nil {
		return err
	}

	groupEntries, err := groupSheets(sheets)
	if err != nil {
		return nil
	}

	return outputGroupEntries(groupEntries)
}
