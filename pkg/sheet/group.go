package sheet

func groupSheets(sheets []Sheet) ([]GroupEntry, error) {
	g := make([]GroupEntry, len(sheets))

	for idx, sheet := range sheets {
		entries := groupEntries(sheet.Entries)

		g[idx] = GroupEntry{
			Name:    sheet.Name,
			Entries: entries,
		}
	}

	return g, nil
}

func groupEntries(d []Entry) map[string]WKNSumEntry {
	entries := make(map[string]WKNSumEntry)

	for _, entry := range d {
		e, ok := entries[entry.WKN]
		if ok {
			e.Paid += entry.Paid
			e.Costs += entry.Costs
			e.Amount += entry.Amount

			entries[entry.WKN] = e
		} else {
			entries[entry.WKN] = WKNSumEntry{
				Paid:   entry.Paid,
				Costs:  entry.Costs,
				Amount: entry.Amount,
			}
		}
	}

	return entries
}
