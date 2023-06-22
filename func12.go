package main

import (
	"errors"
	"unicode"
)

func FiltrarStringsMaiusculas(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", errors.New("slice vazio")
	}

	var filteredStrings []string

	for _, str := range strings {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			filteredStrings = append(filteredStrings, str)
		}
	}

	result := strings.Join(filteredStrings, " ")

	return result, nil
}
func main() {
	// Exemplo de uso
	slice := []string{"Abacaxi", "banana", "Cenoura", "Damasco"}
	filtered, err := FiltrarStringsMaiusculas(slice)

	if err != nil {
		// Lidar com o erro
	} else {
		// Usar a string filtrada
	}
}
