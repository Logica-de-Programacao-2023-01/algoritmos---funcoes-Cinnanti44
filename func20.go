package main

import (
	"errors"
)

func filterStrings(stringsSlice []string) ([]string, error) {
	if len(stringsSlice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	filteredStrings := make([]string, 0)

	for _, str := range stringsSlice {
		if len(str) > 5 {
			filteredStrings = append(filteredStrings, str)
		}
	}

	return filteredStrings, nil
}
