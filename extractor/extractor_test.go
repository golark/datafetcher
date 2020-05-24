package extractor_test

import (
	"github.com/golark/datagrabber/extractor"
	"testing"
)


const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func init() {
	//log.SetLevel(log.PanicLevel) // do not log during testing below panic
}

func TestDownloadLink(t *testing.T) {

	// test 1 - download link
	testUrl := "https://yahoo.com"
	t.Logf("Test 1:\twhen trying to download url %v, checking for nil error", testUrl)
	_, err := extractor.DownloadLink(testUrl)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)

}

func TestExtractTableFromUrl(t *testing.T) {


	// test 1 - url for extraction
	testUrl := "https://drive.google.com/open?id=1KHNY9xA230doAgD-5h3TMRo0P-SlptiA"
	t.Logf("Test 1:\twhen trying to extract table from url %v, checking for nil error", testUrl)

	tableWeb, err := extractor.ExtractTableFromUrl(testUrl)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)
	if tableWeb.Data == nil {
		t.Fatalf("\t%s\tshould return data", failed)
	}
	t.Logf("\t%s\tshould return data", succeed)


	// test 2 - extract from local csv
	t.Logf("Test 2:\twhen trying to read local csv file, checking for nil error")
	tableLocal, err := extractor.ExtractTableFromFile("../testfiles/test.csv")
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)


	// test 3 - compare local and web table extraction
	t.Logf("Test 3:\twhen trying to compare local and web file table extraction, checking for bit exactness")
	for i, k := range tableLocal.RowHeaders {
		if k != tableWeb.RowHeaders[i] {
			t.Fatalf("\t%s\tmismatch in row headers %v", failed, k)
		}
	}
	t.Logf("\t%s\tshould match row headers", succeed)

	for i, k := range tableLocal.ColHeaders {
		if k != tableWeb.ColHeaders[i] {
			t.Fatalf("\t%s\tmismatch in col headers %v", failed, k)
		}
	}
	t.Logf("\t%s\tshould match col headers", succeed)
}