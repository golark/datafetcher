package extractor

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)


// DownloadLink
// Downloads the given URL and saves the contents to a temporary file
func DownloadLink(URL string) ([][]string, error) {

	// step 1 - download the file
	resp, err := http.Get(URL)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "url": URL}).Error("cant get URL")
		return nil, err
	}
	defer resp.Body.Close()


	// step 2 - read it all
	contents, err := readCsvContents(resp.Body)

	return contents, nil
}

// DownloadCsvFile
// download csv file from URL and save to filePath
func DownloadCsvFile(url string, filePath string, timeoutSeconds int) ([][]string, error) {

	// step 1 - get http
	client := http.Client{Timeout:time.Duration(timeoutSeconds) * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http request status is not OK")
	}

	// step 2 - check file type from header
	contentType := resp.Header["Content-Type"][0]
	if strings.Contains(contentType, "text/csv") == false {
		log.WithFields(log.Fields{"err:": err, "file type": resp.Header["Content-Type"][0]}).Error("")
		return nil, errors.New("downloaded content type is not text/csv")
	}

	// step 2 - create a file
	f, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// step 3 - write to file
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return nil, err
	}

	// step 4 - open file and read contents
	contents, err := readCsvContents(f)

	return contents, nil
}

