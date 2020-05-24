package symphoniser_test

import (
	"github.com/golark/datagrabber/symphoniser"
	"github.com/labstack/gommon/random"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func TestImportTableToDb(t *testing.T) {

	t.Logf("Test 1:\twhen trying to import table to dB, checking for nil error")
	identifier := random.String(10)
	data := [][]string {
		[]string{"11", "12", "13", "14", "15", "16"},
		[]string{"21", "22", "23", "24", "25", "26"},
		[]string{"31", "32", "33", "34", "35", "36"},
		[]string{"41", "42", "43", "44", "45", "46"},
		[]string{"51", "52", "53", "54", "55", "56"},
	}
	rowHead := []string { "r1", "r2", "r3", "r4", "r5"}
	colHead := []string { "c1", "c2", "c3", "c4", "c5", "c6"}

	err := symphoniser.ImportTableToDb(data, rowHead, colHead, identifier)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)

}
