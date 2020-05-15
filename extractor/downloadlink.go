package extractor

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

// DownloadLink
// Downloads the given URL and saves the contents to a temporary file
func DownloadLink(URL string) ([][]string, error) {

	log.WithFields(log.Fields{"url": URL}).Info("downloading link")

	// step 1 - download the file
	resp, err := http.Get(URL)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "url": URL}).Error("cant get URL")
		return nil, err
	}
	defer resp.Body.Close()

	// step 2 - read it all
	contents, err := ReadFileContents(resp.Body)

	return contents, nil
}
