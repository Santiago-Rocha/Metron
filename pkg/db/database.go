package db

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../conf"
)

type MongoDatastore struct {
	db      *mongo.Database
	Session *mongo.Client
}

func (dataStore *MongoDatastore) GetCollection(collection string) *mongo.Collection {
	return dataStore.db.Collection(collection)
}

func NewDatastore() *MongoDatastore {
	var config = conf.GetConfiguration()
	var mongoDataStore *MongoDatastore
	db, session := connect(config)
	if db != nil && session != nil {

		// log statements here as well

		mongoDataStore = new(MongoDatastore)
		mongoDataStore.db = db
		mongoDataStore.Session = session
		return mongoDataStore
	}

	logrus.Fatalf("Failed to connect to database: %v", config.Database.Name)

	return nil
}

func connect(config conf.Configuration) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(config)
	})

	return db, session
}

func connectToMongo(config conf.Configuration) (a *mongo.Database, b *mongo.Client) {

	var err error
	logrus.Info("Creating Session")
	session, err := mongo.NewClient(options.Client().ApplyURI(config.Database.ConnectionUri))
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Creating Connection")
	err = session.Connect(context.TODO())
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Testing Connection")
	err = session.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}

	var DB = session.Database(config.Database.Name)
	logrus.Infof("Successfully connected to database: %v", config.Database.Name)

	return DB, session
}
