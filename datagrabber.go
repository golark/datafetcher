package main

import (
	"github.com/golark/datagrabber/explorer"
	log "github.com/sirupsen/logrus"
)

func main() {

	url:= "https://data.humdata.org/dataset/novel-coronavirus-2019-ncov-cases"
	linkTraces := explorer.FindLinksOnPage(url)

	filtTraces := explorer.FilterLinkTraces(linkTraces, []string{".csv"})
	filtTraces  = explorer.FilterLinkTraces(filtTraces, []string{"covid", "corona"})

	for _, trace := range filtTraces {
		log.WithFields(log.Fields{"trace Text:":trace.Text}).Info("")
		log.WithFields(log.Fields{"trace URL:":trace.Url}).Info("")
	}

}
