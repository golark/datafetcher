package classifier

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestGetRuneCategories(t *testing.T) {

	testString := "Hello World"
	expectedRuneCats := []string {"Zs", "Ll", "Lu", "Z", "L"}


	// perform test
	runeCategories := getRuneCategories(testString)


	// check the results
	// criteria: rune categories must be one of the expectedRuneCategories above
	log.WithFields(log.Fields{"runeCategories":runeCategories}).Info("")
	for _, cat := range runeCategories {

		match := false
		for _, expCat := range expectedRuneCats {
			if expCat == cat {
				match = true
				break
			}
		}
		if !(match) {
			t.Error("detected rune category did not match one from expected rune categories")
		}


	}


}
