package db_test

import (
	"github.com/golark/datagrabber/db"
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

func TestAddCollection(t *testing.T) {

	// test 1
	t.Logf("Test 1:\twhen trying to connect to db at %v, checking for nil error",URI)

	client, err := db.Connect(URI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	defer db.Disconnect(client)
	t.Logf("\t%s\tshould return nil err", succeed)


	// test 2
	database := "testdb"
	collectionURI := "testcollection"
	t.Logf("Test 2:\twhen trying to add collection: %v and databse: %v, checking for nil error",database, collectionURI)

	_, err = db.AddCollection(client, database, collectionURI)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)
}


