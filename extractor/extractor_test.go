package extractor_test

import (
	"github.com/golark/datagrabber/extractor"
	log "github.com/sirupsen/logrus"
	"testing"
)


func init() {
	log.SetLevel(log.PanicLevel) // do not log during testing below panic
}

func TestDownloadLink(t *testing.T) {

	testLink := "https://yahoo.com"
	_, err := extractor.DownloadLink(testLink)
	if err != nil {
		t.Error("Cant download link")
	}
}

