package classifier

import "unicode"


// classifierFunc
// determines if the string belongs to a given class
type classifierFunc func(string) bool

const (
	NoClass     = iota
	DateClass   = iota
	NumberClass = iota
	TimeClass   = iota
	LetterClass = iota
)


// Classes
// classes and the classification functions that determine whether given string belongs to one of the below classes
var Classes = map[int]classifierFunc {
	DateClass : isDate,
	NumberClass : isNumerical,
	TimeClass : isTime,
	LetterClass : isLetters,
}

// isTime
// returns true if only numeric and ':' runes constitute input string
// there must be at least one divider (  ':' )
func isTime(s string) bool {

	if len(s) < 4 { // check have time with 4 or less runes
		return false
	}

	consecutiveNums := 0
	dividerCount := 0
	for _, r := range s {
		if unicode.IsNumber(r) {
			consecutiveNums++
			if consecutiveNums > 3 { // cant have more than 3 consecutive numbers in time unless some atomic time
				return false
			}
		} else if r == ':' {
			consecutiveNums = 0
			dividerCount++
		} else if !unicode.IsSpace(r) {
			return false
		}
	}

	if dividerCount < 1 { // there must be at least one divider in time
		return false
	}

	return true
}

// isLetters
// returns true if the data only consists of letters, space and punctuatio ( no numeric digits are allowed )
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

	if numDividers == 0 || numDividers > 2 || consecutiveNums > 4{
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
// try to classify input string to one of the "Classes"
func Classify(s string) (int){

	for class, classFunc := range Classes {
		if classFunc(s) {
			return class
		}
	}

	return NoClass // can't classify

}

