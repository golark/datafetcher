package extractor

import "testing"

func TestDownloadLink(t *testing.T) {

	testLink := "https://yahoo.com"
	_, err := DownloadLink(testLink)
	if err != nil {
		t.Error("Cant download link")
	}
}

