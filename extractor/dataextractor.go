package extractor

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"io"
)

// ReadFileContents
// read file contents
func ReadFileContents(r io.ReadCloser) ([][]string, error ){

	reader := csv.NewReader(r)
	contents, err := reader.ReadAll()
	if err!=nil {
		log.WithFields(log.Fields{"err":err}).Error("cant read reader contents")
		return nil, nil
	}
	if contents != nil { // log row/col count
		log.WithFields(log.Fields{"cols": len(contents[0]), "rows:":len(contents)}).Info("num columns")
	}

	return contents, nil
}

// extractHeaders
// return headers ( first column and first row )
func extractHeaders(data [][]string) (rowHeaders, colHeaders []string){

	rowHeaders = data[0]

	colHeaders = make([]string, len(data))

	for i:=0;i<len(data);i++ {
		colHeaders[i] = data[i][0]
	}

	return rowHeaders, colHeaders

}

func GetDataHeadersFromUrl(url string) (rowHeaders, colHeaders []string){

	linkContents, err := DownloadLink(url)
	if err!= nil {
		log.WithFields(log.Fields{"err":err}).Error("cant download link")
		return nil, nil
	}

	rowHeaders, colHeaders = extractHeaders(linkContents)

	return rowHeaders, colHeaders
}

