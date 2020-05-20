package classifier_test

import (
	"github.com/golark/datagrabber/classifier"
	log "github.com/sirupsen/logrus"
	"testing"
)


const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func init() {
	log.SetLevel(log.PanicLevel) // do not log during testing below panic
}

func TestClassifyTable(t *testing.T) {

	testCases := map[string]int{
		"15:30": classifier.TimeClass,
		"aveebeesfs": classifier.LetterClass,
		"2/3/2015": classifier.DateClass,
		"avvsasa?1223121wff...": classifier.NoClass,
		"2.1.2.2.3": classifier.NumberClass,
	}

	textIdx := 0
	for s,r := range testCases {
		t.Logf("Test %v:\twhen trying to classify %v, checking for class %v",textIdx, s, r)

		res := classifier.Classify(s)
		if res != r {
			t.Fatalf("\t%s\tshould return %v", failed, r)
		}
		t.Logf("\t%s\tshould return %v", succeed, r)

		textIdx++
	}
}
