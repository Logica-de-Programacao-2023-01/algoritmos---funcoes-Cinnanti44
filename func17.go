package main

import (
	"errors"
	"sort"
	"strings"
)

func concatenateStrings(stringsSlice []string) (string, error) {
	if len(stringsSlice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	// Ordena o slice de strings em ordem alfabética
	sort.Strings(stringsSlice)

	// Concatena as strings ordenadas separadas por espaço
	result := strings.Join(stringsSlice, " ")

	return result, nil
}
