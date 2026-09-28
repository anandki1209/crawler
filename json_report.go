package main

import (
	"encoding/json"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sorted []PageData
	for _, k := range keys {
		sorted = append(sorted, pages[k])
	}
	data, err := json.MarshalIndent(sorted, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(filename, data, 0644)

	return nil

}
