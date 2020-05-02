package explorer

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// LinkTrace
// text/url tuple that are extracted from pages
type LinkTrace struct{
	DataIdentifier string
	Url string
	PrunedDataIdentifier string
}

// FindLinksOnPage
// find all the links on the page
func FindLinksOnPage(url string) []LinkTrace {

	var linkTraces []LinkTrace

	// step 1 - Instantiate default collector
	c := colly.NewCollector()

	// step 2 - append results to collection during search
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		linkTraces = append(linkTraces, LinkTrace{e.Text, e.Attr("href"), ""})
	})

	// step 3 - before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.WithFields(log.Fields{"url":r.URL.String()}).Info("searching page")
	})

	// step 4 - start scraping
	c.Visit(url)

	// step 5 - warn if we can not find any links
	if linkTraces == nil {
		log.WithFields(log.Fields{"url":url}).Warn("could not find any links on page")
	}

	return linkTraces

}

// FilterLinkTraces
// look for the given filters in the linkTraces text and url
// any hit on any of the filters are registered
// returns a filtered slice
// input parameters as well as the filter are converted to lower case prior to search
func FilterLinkTraces(linkTraces []LinkTrace, filters []string) []LinkTrace {

	var  filteredTraces []LinkTrace

	// convert filters to lower case
	var lcFilter []string
	for _, filter := range filters {
		lcFilter = append(lcFilter, strings.ToLower(filter))
	}

	// check both trace text and url for the filter
	for _, trace := range linkTraces {

		// match case to the filter
		text := strings.ToLower(trace.DataIdentifier)
		url  := strings.ToLower(trace.Url)

		for _, filter := range lcFilter {
			if (strings.Contains(text, filter)) ||  (strings.Contains(url, filter)) {
				filteredTraces = append(filteredTraces, trace)
				break
			}
		}
	}

	return filteredTraces
}

// PruneDataIdentifier
// search through link trace text to match following rule and populate PrunedDataIdentifier of the trace
// any non space character, followed by dataIdentifier, followed by any of the extensions
// case insensitive
func PruneDataIdentifier(linkTraces []LinkTrace, dataIdentifier string) {

	// any non space character, followed by dataIdentifier, followed by any of the extensions
	re := regexp.MustCompile(`\S*` + strings.ToLower(dataIdentifier) + `.*(.csv|.doc|csv|.json)`)

	for i, trace := range linkTraces {
		text := strings.ToLower(trace.DataIdentifier)
		linkTraces[i].PrunedDataIdentifier = re.FindString(text)
	}

}
