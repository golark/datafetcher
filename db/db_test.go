package db_test

import (
	"github.com/golark/datagrabber/db"
	"github.com/labstack/gommon/random"
	log "github.com/sirupsen/logrus"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"

	URI        = "mongodb://localhost:27017"
)

func init() {
	log.SetLevel(log.PanicLevel) // do not log during testing below panic
}

func TestConnectDisconnect(t *testing.T) {

	t.Logf("Test:\twhen trying to connect to db at %v, checking for nil error",URI)

	client, err := db.Connect(URI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	defer db.Disconnect(client)

	t.Logf("\t%s\tshould return nil err", succeed)
}

func TestInsertSingleDataPoint(t *testing.T) {

	// test 1 - connect to db
	t.Logf("Test 1:\twhen trying to connect to db at %v, checking for nil error",URI)

	client, err := db.Connect(URI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)


	// test 2 - add collection
	database := "testdb"
	collectionURI := "testcollection"
	t.Logf("Test 2:\twhen trying to add collection: %v and databse: %v, checking for nil error",database, collectionURI)

	collection, err := db.AddCollection(client, database, collectionURI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)


	// test 3 - add document
	t.Logf("Test 3:\twhen trying to add a data point to collection %v, checking for nil error",collectionURI)

	row := random.String(5)
	col := random.String(5)
	val := random.String(5)
	db.InsertSingleDataPoint(collection, db.DataPoint{Col:col, Row:row, Val:val})
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)


	// test 4 - find document
	t.Logf("Test 4:\twhen trying to fetch a data point from collection %v, checking for val: %v",collectionURI, val)

	dp, err := db.GetSingleDataPoint(collection, row, col)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	if dp.Val != val {
		t.Fatalf("\t%s\tshould return val %v", failed, val)
	}
	t.Logf("\t%s\tshould return %v", succeed, val)


	// test 5 - remove collection
	t.Logf("Test 5:\twhen trying to remove collection %v, checking for nil error",collectionURI)

	err = db.RemoveCollection(client, database, collectionURI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)


	// test 6 - disconnect
	t.Logf("Test 6:\twhen trying to disconnect db client checking for nil error")

	err  = db.Disconnect(client)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)

}

