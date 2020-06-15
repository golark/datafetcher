package detective

import (
	"errors"
	"github.com/golark/datagrabber/db"
	log "github.com/sirupsen/logrus"
)

var (
	ErrNoCollectionMatch = errors.New("no matching collection in database")
	ErrNoIdentifierMatch = errors.New("no matching identifier in collection")
)

// SearchDatabase
// first try to find the collection in db
// if collection exists try to find the data
func SearchDatabase(URIdB string, databaseName string, collectionIdentifier string, dataIdentifier string) (db.Line, error) {

	// step 1 - get collection names from db
	client, err := db.Connect(URIdB)
	if err != nil {
		return db.Line{}, err
	}
	collectionNames, err := db.GetCollectionNames(client, databaseName)
	if err != nil {
		return db.Line{}, err
	}

	// step 2 - search through collection names
	// looking for bit-exact match
	collectionFound := false
	for _, s := range collectionNames {
		if collectionIdentifier == s {
			collectionFound = true
		}
	}
	if collectionFound == false {
		return db.Line{}, ErrNoCollectionMatch
	}

	// step 3 - search inside collection for data
	// get document identifiers
	collection, err := db.GetCollection(client, databaseName, collectionIdentifier)
	if err != nil {
		return db.Line{}, err
	}
	identifiers, err := db.GetIdentifiers(collection)

	log.WithFields(log.Fields{"identifiers": identifiers}).Info("")
	if err != nil {
		return db.Line{}, err
	}
	identifierFound := false
	for _, i := range identifiers {
		if i == dataIdentifier {
			identifierFound = true
			break
		}
	}
	if identifierFound == false {
		return db.Line{}, ErrNoIdentifierMatch
	}

	// step 4 - get the line corresponding to the detected identifier in detected collection
	l, err := db.GetSingleLine(collection, dataIdentifier)

	return l, nil
}