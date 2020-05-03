package classifier

import "unicode"

// isNumerical
// returns true if the data only consists of numbers and dots
func isNumerical(s string) bool {

	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

// getRuneCategories
// returns unique rune categories that the input string constitutes
func getRuneCategories(s string) []string {

	var uniqueCats []string // unique categories
	for category, t := range unicode.Categories {

		for _, r := range s {
			if unicode.Is(t, r) {
				uniqueCats = append(uniqueCats, category)
			}
		}
	}

	return uniqueCats
}

// Classify
// classify data as one of the following:
//
func Classify(identifier string) (string, int){
	return "", 0
}
