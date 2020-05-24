package db

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrClientIsNotInitialised = errors.New("client is not initialised")
	ErrDatabaseDoesntExist    = errors.New("database doesnt exist")
	ErrCantConnect            = errors.New("cant connect to db")
	ErrCantDisconnect         = errors.New("cant terminate db connection")
	ErrCollectioniIsNil       = errors.New("collection is nil")
)

type Line struct {
	Identifier string
	X []string
	Y []string
}

// InsertSingleDataPoint
func InsertSingleLine(collection *mongo.Collection, l Line) error {

	// step 1 - insert single line
	res, err := collection.InsertOne(context.TODO(), l)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"insertedId":res.InsertedID}).Trace("inserted item")

	return nil
}

// GetSingleLine
// search for single line and return match
func GetSingleLine(collection *mongo.Collection, identifier string) (Line, error) {

	var l Line
	// step 1 - check collection
	if collection == nil {
		return l, ErrCollectioniIsNil
	}

	// step 2 - find one and decode into a datapoint
	filter := bson.D{{"identifier", identifier }}
	err := collection.FindOne(context.TODO(), filter).Decode(&l)

	return l, err
}

// DataPoint
// a single entry of data with row/col headers
type DataPoint struct {
	Row string
	Col string
	Val string
}

// InsertSingleDataPoint
func InsertSingleDataPoint(collection *mongo.Collection, dp DataPoint) error {

	// step 1 - insert single data point
	res, err := collection.InsertOne(context.TODO(), dp)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"insertedId":res.InsertedID}).Trace("inserted item")

	return nil
}

// GetSingleDataPoint
// search collection for single match with row/col parameters and decode into a datapoint
func GetSingleDataPoint(collection *mongo.Collection, row string, col string) (DataPoint, error) {

	var dp DataPoint

	// step 1 - check collection
	if collection == nil {
		return dp, ErrCollectioniIsNil
	}

	// step 2 - find one and decode into a datapoint
	filter := bson.D{{"row", row }, {"col", col}}
	err := collection.FindOne(context.TODO(), filter).Decode(&dp)

	return dp, err
}

// InsertTable
// insert double dimensional slice of datapoints
func InsertTable(collection *mongo.Collection, table [][]DataPoint) {

	for _,i := range table {
		InsertDataPointLine(collection, i)
	}

}

// InsertLine
// insert slice of data points
func InsertDataPointLine(collection *mongo.Collection, line []DataPoint) error {
	for _,i := range line {
		err := InsertSingleDataPoint(collection, i)
		if err!= nil {
			return err
		}
	}
	return nil
}


// IsConnected
// returns nil if can ping db
func IsConnected(client *mongo.Client) error {

	// step 1 - check if client is nil
	if client == nil {
		return ErrClientIsNotInitialised
	}

	// step 2 - try to ping
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("cant ping db")
		return ErrCantConnect
	}
	log.Trace("successfully pinged mongodb")

	return nil
}

// Connect
// connect to db
func Connect(uri string) (*mongo.Client, error) {

	// step 1 - check URI
	if uri == "" {
		log.Error("URI can not be empty")
		return nil, errors.New("URI can not be empty")
	}

	// step 2 - connect to mongodb
	log.WithFields(log.Fields{"uri": uri}).Trace("attempting connection to db")

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "URI": uri}).Error("error while trying to connect to db")
		return nil, ErrCantConnect
	}

	// step 3 - sanity check
	if err := IsConnected(client); err != nil {
		return nil, err
	}
	log.Trace("connected to mongodb")

	return client, nil
}

// Disconnect
// terminate db connection
func Disconnect(client *mongo.Client) error {

	// step 1 - first check if we have a valid client
	if err := IsConnected(client); err != nil {
		return err
	}

	// step 2 - terminate
	if err := client.Disconnect(context.TODO()); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("error while disconnecting from db")
		return ErrCantDisconnect
	}

	log.Trace("terminated db connection")

	return nil
}

// GetCollection
// new database and collection unless exits
func GetCollection(client *mongo.Client, database string, collection string) (*mongo.Collection, error) {

	// step 1 - check client connection
	if err := IsConnected(client); err != nil {
		return nil, err
	}

	// step 2 - add database and collection
	return client.Database(database).Collection(collection), nil
}

// RemoveCollection
func RemoveCollection(client *mongo.Client,  database string, collection string) error {

	// step 1 - check client connection
	if err := IsConnected(client); err != nil {
		return err
	}

	// step 2 - remove collection
	err := client.Database(database).Collection(collection).Drop(context.TODO())

	return err
}

