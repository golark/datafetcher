package classifier

import "unicode"

var Classes = map[string][]*unicode.RangeTable {
	// "date"     : {unicode.N},
	"integer"  : {unicode.N},
	// "floating" : {unicode.N, },
	"time"     : {unicode.N},
	"currency" : {unicode.N, unicode.Sc},
	"letters"  : {unicode.Lu, unicode.Ll, unicode.L, unicode.Z, unicode.P},
}

// isLetters
// returns true if the data only consists of letters and punctuation
func isLetters(s string) bool {

	if s == "" { // check empty string
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsPunct(r) && !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}


// isDate
func isDate(s string) bool {

	// step 1 -  check minimum rune requirement for a date
	if len(s) < 4 {
		return false
	}

	// step 2 -
	// - there must be at least one instance of a '.' or '/' but less than 2 dividers
	// - others must be a numerical digit, and no more than 4 consecutive numbers
	numDividers := 0
	consecutiveNums := 0
	for _, r := range s {

		if unicode.IsNumber(r) {
			consecutiveNums++
			if consecutiveNums > 4 { // no more than 4 consecutive numbers allowed
				return false
			}
		} else if r == '/' || r == '.' {
			numDividers++
			consecutiveNums = 0
		} else if !unicode.IsSpace(r) {
			return false // if not space or number or '/' or '.' this cant be a date
		}
	}

	if numDividers == 0 || numDividers > 2 {
		return false
	} else {
		return true
	}

}

// isNumerical
// returns true if the data only consists of numbers and dots
func isNumerical(s string) bool {

	if s == "" { // check empty string
		return false
	}

	for _, r := range s {
		if !unicode.IsNumber(r) && r != '.' {
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

