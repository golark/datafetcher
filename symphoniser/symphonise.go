package symphoniser

import (
	"github.com/golark/datagrabber/explorer"
	"github.com/golark/datagrabber/extractor"
	log "github.com/sirupsen/logrus"
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

