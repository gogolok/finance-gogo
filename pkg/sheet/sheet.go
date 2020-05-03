package sheet

func Data() error {
	config, err := ParseConfig()
	if err != nil {
		return err
	}

	sourceRawData, err := fetchAllSourceRawData(config.Sources)
	if err != nil {
		return err
	}

	sheets, err := parseAll(sourceRawData)
	if err != nil {
		return err
	}

	return outputSheets(sheets)
}
