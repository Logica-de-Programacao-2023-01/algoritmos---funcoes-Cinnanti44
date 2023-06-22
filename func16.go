package main

import (
	"errors"
	"strings"
)

func replaceVowels(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("A string está vazia")
	}

	vowels := "aeiouAEIOU"
	replacedStr := ""

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			replacedStr += "*"
		} else {
			replacedStr += string(char)
		}
	}

	return replacedStr, nil
}
