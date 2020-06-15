package symphoniser_test

import (
	"github.com/golark/datagrabber/extractor"
	"github.com/golark/datagrabber/symphoniser"
	"github.com/labstack/gommon/random"
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

func TestImportTableTodB(t *testing.T) {

	// test 1 - import table to db
	t.Logf("Test 1:\twhen trying to import table to dB, checking for nil error")
	collectionURI := random.String(10)
	data := [][]string{
		[]string{"11", "12", "13", "14", "15", "16"},
		[]string{"21", "22", "23", "24", "25", "26"},
		[]string{"31", "32", "33", "34", "35", "36"},
		[]string{"41", "42", "43", "44", "45", "46"},
		[]string{"51", "52", "53", "54", "55", "56"},
	}
	rowHead := []string{"r1", "r2", "r3", "r4", "r5"}
	colHead := []string{"c1", "c2", "c3", "c4", "c5", "c6"}

	err := symphoniser.ImportTableTodB(data, rowHead, colHead, collectionURI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)

	// test 2 - export line from db and compare
	t.Logf("Test 2:\twhen trying to export line from dB, checking for nil error and line match")
	l, err := symphoniser.ExportLine(collectionURI, rowHead[3])
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)
	if l.Identifier != rowHead[3] {
		t.Fatalf("\t%s\tline identifier mismatch %v", failed, l.Identifier)
	}
	t.Logf("\t%s\tshould match line identifier", succeed)
	for i, d := range l.Y {
		if d != data[3][i] {
			t.Fatalf("\t%s\tline data mismatch %v", failed, d)
		}
	}
	t.Logf("\t%s\tshould match line data", succeed)

}

func TestImportCsvTodB(t *testing.T) {

	// test 1 - read local csv
	t.Logf("Test 1:\twhen trying to read local csv file, checking for nil error")
	tableData, err := extractor.ExtractTableFromFile("./test.csv")
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)

	// test 2 - import table to db
	t.Logf("Test 2:\twhen trying to import table to dB, checking for nil error")
	collectionURI := random.String(10)

	err = symphoniser.ImportTableTodB(tableData.Data, tableData.RowHeaders, tableData.ColHeaders, collectionURI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)

	// test 3 - export line from db and compare
	t.Logf("Test 3:\twhen trying to export line from dB, checking for nil error and line match")
	l, err := symphoniser.ExportLine(collectionURI, tableData.RowHeaders[0])
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)
	if l.Identifier != tableData.RowHeaders[0] {
		t.Fatalf("\t%s\tline identifier mismatch %v", failed, l.Identifier)
	}
	t.Logf("\t%s\tshould match line identifier", succeed)
	for i, d := range l.Y {
		if d != tableData.Data[0][i] {
			t.Fatalf("\t%s\tline data mismatch %v", failed, d)
		}
	}
	t.Logf("\t%s\tshould match line data", succeed)

}

func TestDataInquiry(t *testing.T) {

	// prep 1 - read local csv
	tableData, err := extractor.ExtractTableFromFile("./test.csv")
	if err != nil {
		t.Fatalf("\t%s\tprep 1 - should not return %v", failed, err)
	}

	// prep 2 - import table to db
	collectionURI := "testcollection"
	err = symphoniser.ImportTableTodB(tableData.Data, tableData.RowHeaders, tableData.ColHeaders, collectionURI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}

	// test 1 - inquire data
	dataIdentifier := collectionURI + " " + tableData.RowHeaders[0]
	t.Logf("Test 1:\twhen inquiring data %v, checking for nil error", dataIdentifier)
	_, err = symphoniser.DataInquiry(dataIdentifier)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)

}
