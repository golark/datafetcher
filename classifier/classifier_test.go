package classifier

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestClassify(t *testing.T) {

	// test 1
	testString := "15:30"
	exp := TimeClass
	res := Classify(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 2
	testString = "aveebeesfs"
	exp = LetterClass
	res = Classify(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 3
	testString = "2/3/2015"
	exp = DateClass	// test 4
	testString = "avvsa?132135wf..."
	exp = NoClass
	res = Classify(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}
	res = Classify(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 4
	testString = "avvsa?132135wf..."
	exp = NoClass
	res = Classify(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 5
	testString = "2.1.232015"
	exp = NumberClass
	res = Classify(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

}

func TestIsTime(t *testing.T) {

	// test 1
	testString := "15:30"
	exp := true
	res := isTime(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 2
	testString = "1649"
	exp = false
	res = isTime(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 3
	testString = "abcv"
	exp = false
	res = isTime(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}
}

func TestIsLetters(t *testing.T) {

	// test 1
	testString := "88"
	exp := false
	res := isLetters(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 2
	testString = "acbvAdaufsa.fa:"
	exp = true
	res = isLetters(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 3
	testString = "acbvAdaufsa.....? fa:"
	exp = true
	res = isLetters(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}
}

func TestIsNumerical(t *testing.T) {

	// test 1
	testString := "22"
	exp := true
	res := isNumerical(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 2
	testString = "22aa"
	exp = false
	res = isNumerical(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 3
	testString = "//11223//"
	exp = false
	res = isNumerical(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 4
	testString = "11.22"
	exp = true
	res = isNumerical(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 6
	testString = "13:24"
	exp = false
	res = isNumerical(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}
}


func TestIsDate(t *testing.T) {

	// test 1
	testString := "20/2020"
	exp := true
	res := isDate(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 2
	testString = "1/01/1998"
	exp = true
	res = isDate(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 3
	testString = "1-asad"
	exp = false
	res = isDate(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 4
	testString = "1152421"
	exp = false
	res = isDate(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 5
	testString = "11.05.2020"
	exp = true
	res = isDate(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}

	// test 6
	testString = "11/22/2222/222/22//"
	exp = false
	res = isDate(testString)
	if res != exp {
		t.Errorf("expected %v got %v for testString %v", exp, res, testString)
	}
}

func TestGetRuneCategories(t *testing.T) {

	testString := "Hello World"
	expectedRuneCats := []string {"Zs", "Ll", "Lu", "Z", "L"}


	// perform test
	runeCategories := getRuneCategories(testString)


	// check the results
	// criteria: rune categories must be one of the expectedRuneCategories above
	log.WithFields(log.Fields{"runeCategories":runeCategories}).Trace("")
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
