package sheet

import ()

func Data() error {
	config, err := ParseConfig()
	if err != nil {
		return err
	}
	sourceRawData, err := fetchAllSourceRawData(config.Sources)
	if err != nil {
		return err
	}
	return parseAll(sourceRawData)
}
