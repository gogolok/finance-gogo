package sheet

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func fetchAllSourceRawData(sources []SourceUrl) ([]SourceRawData, error) {
	entries := []SourceRawData{}

	for _, source := range sources {
		content, err := DownloadFile(source.Url)
		if err != nil {
			return entries, err
		}

		r := csv.NewReader(strings.NewReader(content))
		r.Comma = ','

		records, err := r.ReadAll()
		if err != nil {
			return entries, err
		}

		d := SourceRawData{
			Name:    source.Name,
			Records: records,
		}

		entries = append(entries, d)

	}

	return entries, nil
}

func DownloadFile(url string) (string, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to download from url %s, status code != 200: %v", url, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
