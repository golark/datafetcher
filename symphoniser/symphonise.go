package symphoniser

import (
	"errors"
	"github.com/golark/datagrabber/db"
	"github.com/golark/datagrabber/dgproto"
	"github.com/golark/datagrabber/explorer"
	"github.com/golark/datagrabber/extractor"
	log "github.com/sirupsen/logrus"
)

const (
	URIdB = "mongodb://localhost:27017"
	DATABASE = "DATAGRABBER"
)

// DataInquiry
// search db for the data and serve directly if data is local
// search web if data is not local
func DataInquiry(identifier string) ([]dgproto.PointResp, error) {

	// @todo: step 1 - search db for the requested data

	// @todo: step 2 - serve data if exists on local db

	// @todo: step 3 - search web if data is not local

	return nil, nil
}

// GetDataHeaders
// return only data headers that might be related to the data identifier
func GetDataHeaders(dataIdentifier string) (rowHeaders, colHeaders []string) {

	// step 1 - find data resource
	// @todo: make the url search based
	url:= "https://data.humdata.org/dataset/novel-coronavirus-2019-ncov-cases"

	// step 2 - find link traces
	linkTraces := explorer.SearchLinkTraces(url, []string{dataIdentifier}, ".csv")

	// step 3 - prune data identifiers
	explorer.PruneDataIdentifier(linkTraces, dataIdentifier)

	// step 4 - get headers
	rowHeaders, colHeaders = extractor.GetDataHeadersFromUrl("https://data.humdata.org/"+linkTraces[0].Url)

	log.WithFields(log.Fields{"rowHeaders":rowHeaders}).Info("")
	log.WithFields(log.Fields{"colHeaders":colHeaders}).Info("")

	return rowHeaders, colHeaders
}

// ImportTableTodB
func ImportTableTodB(data [][]string, rowHead []string, colHead []string, identifier string) error {

	// step 1 - check data size for integrity
	if len(rowHead) != len(data) {
		return errors.New("row size / data mismatch")
	}
	if len(colHead) != len(data[0]) {
		return errors.New("column size / data mismatch")
	}

	// step 2 - connect to db
	client, err := db.Connect(URIdB)
	if err != nil {
		return err
	}

	// step 3 - add collection
	collection, err := db.GetCollection(client, DATABASE, identifier)
	if err != nil {
		return err
	}

	// step 4 - import rows
	for i, r := range rowHead {
		l := db.Line{Identifier:r,
			X: colHead,
			Y: data[i],
		}

		err = db.InsertSingleLine(collection, l)
		if err != nil {
			return err
		}
	}

	// step 5 - import columns
	colData := make([]string, len(data))
	for i, c := range colHead {
		for k, d := range data {
			colData[k] = d[i]
		}

		l := db.Line{Identifier:c,
			X: rowHead,
			Y: colData,
		}

		err = db.InsertSingleLine(collection, l)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExportLine
// get line with the identifier from given collection URI
func ExportLine(collectionURI string, identifier string) (db.Line, error) {

	// step 1 - connect to db
	client, err := db.Connect(URIdB)
	if err != nil {
		return db.Line{}, err
	}

	// step 2 - get collection
	collection, err := db.GetCollection(client, DATABASE, collectionURI)
	if err != nil {
		return db.Line{}, err
	}

	// step 3 - get line from db
	l, err := db.GetSingleLine(collection, identifier)
	if err != nil {
		return db.Line{}, err
	}

	return l, nil
}

