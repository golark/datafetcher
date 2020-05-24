package symphoniser

import (
	"github.com/golark/datagrabber/db"
	"github.com/golark/datagrabber/explorer"
	"github.com/golark/datagrabber/extractor"
	log "github.com/sirupsen/logrus"
)

const (
	URIdB = "mongodb://localhost:27017"
	DATABASE = "DATAGRABBER"
)



func GetDataHeaders(dataIdentifier string) (rowHeaders, colHeaders []string) {

	// step 1 - find data resource
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

func ImportTableTodB(data [][]string, rowHead []string, colHead []string, identifier string) error {

	// step 1 - connect to db
	client, err := db.Connect(URIdB)
	if err != nil {
		return err
	}

	// step 2 - add collection
	collection, err := db.GetCollection(client, DATABASE, identifier)
	if err != nil {
		return err
	}

	// step 3 - import rows
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

	// step 4 - import columns
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

// GetLine
// get line with the identifier from collection
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
