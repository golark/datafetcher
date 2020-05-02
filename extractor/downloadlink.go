package extractor

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// DownloadLink
// Downloads the given URL and saves the contents to a temporary file
func DownloadLink(URL string) ([][]string, error) {

	log.WithFields(log.Fields{"url":URL}).Info("downloading link")

	// step 1 - download the file
	resp, err := http.Get(URL)
	if err != nil {
		log.WithFields(log.Fields{"err":err, "url":URL}).Error("cant get URL")
		return nil, err
	}
	defer resp.Body.Close()

	// step 2 - read it all
	contents, err := ReadFileContents(resp.Body)

	// step 3 - open a temporary file and save the contents
	fl, err := os.OpenFile("/tmp/download.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.WithFields(log.Fields{"err":err}).Error("cant open file, skipping file save")
	} else {
		defer fl.Close()
		// step 4 - dump all to a file
		/*
			_, err = io.Copy(fl, resp.Body)
			if err != nil {
				log.WithFields(log.Fields{"err":err}).Error("cant copy to file")
				return err
			}*/
	}

	return contents, nil
}
