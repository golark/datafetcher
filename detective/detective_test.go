package detective_test

import (
	"github.com/golark/datagrabber/db"
	"github.com/golark/datagrabber/detective"
	"github.com/labstack/gommon/random"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"

	URI = "mongodb://localhost:27017"
)

func init() {
	//log.SetLevel(log.PanicLevel) // do not log during testing below panic
}

func TestSearchDatabase(t *testing.T) {

	// prep 1 - connect to db
	client, err := db.Connect(URI)
	if err != nil {
		t.Fatalf("\t%s\tfailed during prep 1 %v", failed, err)
	}

	// prep 2 - add collections and lines
	database := "testdetective"
	collectionList := []string{random.String(5), random.String(10), random.String(12), random.String(20)}
	identifiersList := []string{random.String(19), random.String(9), random.String(7), random.String(23)}

	for _, c := range collectionList {
		collection, err := db.GetCollection(client, database, c)
		if err != nil {
			t.Fatalf("\t%s\tfailed during test preparation %v", failed, err)
		}

		for _, i := range identifiersList {
			l := db.Line{Identifier:i,
				X: []string{"1", "2", "3", "4", "5", "6"},
				Y: []string{"1", "2", "3", "4", "5", "6"},
			}

			err = db.InsertSingleLine(collection, l)
			if err != nil {
				t.Fatalf("\t%s\tfailed during test preparation %v", failed, err)
			}
		}
	}

	// test 1 - search for collection and identifier
	searchCollection := collectionList[3]
	searchData := identifiersList[2]
	t.Logf("Test 1:\twhen trying to search for collection %v and data %v: checking for nil error and identifier match",searchCollection, searchData)
	l, err := detective.SearchDatabase(URI, database, searchCollection, searchData)
	if err != nil {
		t.Fatalf("\t%s\tshould not return %v", failed, err)
	}
	t.Logf("\t%s\tshould return nil err", succeed)
	if l.Identifier != searchData {
		t.Fatalf("\t%s\tfailed identifier mismatch %v", failed, l.Identifier)
	}
	t.Logf("\t%s\tidentifier should match", succeed)

}

