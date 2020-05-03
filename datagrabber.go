package main

//go:generate protoc datagrabber.proto -I=./dgproto --go_out=plugins=grpc:./dgproto

import (
	"github.com/golark/datagrabber/cmd"
	"github.com/golark/datagrabber/explorer"
	"github.com/golark/datagrabber/extractor"
	log "github.com/sirupsen/logrus"
)

func exampleDataFetch() {

	url:= "https://data.humdata.org/dataset/novel-coronavirus-2019-ncov-cases"
	linkTraces := explorer.FindLinksOnPage(url)

	filtTraces := explorer.FilterLinkTraces(linkTraces, []string{".csv"})
	filtTraces  = explorer.FilterLinkTraces(filtTraces, []string{"covid", "corona"})

	if filtTraces == nil {
		log.Info("no link matching criteria was found")
		return
	}

	// log potential links
	log.WithFields(log.Fields{"num Hits": len(filtTraces)}).Info("number of hits")

	for _, trace := range filtTraces {
		log.WithFields(log.Fields{"trace Text:":trace.DataIdentifier}).Trace("")
		log.WithFields(log.Fields{"trace URL:":trace.Url}).Trace("")
	}

	//
	explorer.PruneDataIdentifier(filtTraces, "covid")

	for _, trace := range filtTraces {
		log.WithFields(log.Fields{"name: ":trace.PrunedDataIdentifier}).Info("potential match")
	}

	// download link
	rowHeaders, colHeaders := extractor.GetDataHeadersFromUrl("https://data.humdata.org/" + filtTraces[0].Url)
	log.WithFields(log.Fields{"rowHeaders":rowHeaders}).Info("")
	log.WithFields(log.Fields{"colHeaders":colHeaders}).Info("")

}

func main() {
	cmd.ServeGrpc()
}

