package main

//go:generate protoc datagrabber.proto -I=./dgproto --go_out=plugins=grpc:./dgproto

import (
	"context"
	"github.com/golark/datagrabber/db"
	"github.com/golark/datagrabber/dgproto"
	"github.com/golark/datagrabber/explorer"
	"github.com/golark/datagrabber/extractor"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
)
// exampleSingleShotGrpcClient
func exampleSingleShotGrpcClient(addr string) error {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:8090",opts...)
	if err != nil {
		log.WithFields(log.Fields{"err":err}).Error("cant dial grpc port")
		return err
	}
	defer conn.Close()

	client := dgproto.NewDataServiceClient(conn)

	r, err := client.DataInquiry(context.Background(), &dgproto.SearchReq{Identifier:""})
	if err != nil {
		log.WithFields(log.Fields{"err":err}).Error("cant get data inquiry client")
		return err
	} else {
		for {
			header, err := r.Recv()
			if err == io.EOF {
				break
			} else {
				log.WithFields(log.Fields{"header:":header}).Info("received new header")
			}
		}
	}
	return nil

}

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


func exampledb() {

	URI := "mongodb://localhost:27017"
	client, err := db.Connect(URI)
	if err != nil {
		log.WithFields(log.Fields{"err":err}).Error("")
	}
	defer db.Disconnect(client)


	database := "testdbt"
	collectionURI := "testcollection"

	collection, err := db.AddCollection(client, database, collectionURI)
	if err != nil {
		log.WithFields(log.Fields{"err":err}).Error("")
	}

	db.InsertSingleDataPoint(collection, db.DataPoint{Col:"col", Row:"row", Val:"val"})
	if err != nil {
	}
}

func main() {
	exampledb()
	// cmd.ServeGrpc()
}

