package sheet

import (
	"time"
)

type Configuration struct {
	Sources []SourceUrl
}

type SourceUrl struct {
	Name string
	Url  string
}

type SourceRawData struct {
	Name    string
	Records [][]string
}

type Entry struct {
	Source string // trace, comdirect, ...
	WKN    string // German securities identification code
	Date   time.Time
	Paid   float64
	Costs  float64 // provision + clearstream + market + whatever costs...
	Amount float64
}

type Sheet struct {
	Name    string
	Entries []Entry
}

type WKNSumEntry struct {
	Paid   float64
	Costs  float64
	Amount float64
}

type GroupEntry struct {
	Name    string
	Entries map[string]WKNSumEntry
}
