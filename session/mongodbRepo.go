package session

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	sessionDbName         = "training_db"
	sessionCollectionName = "sessions"
)

type MongoDBSessionRepository struct {
	db    *mgo.Database
	dbUrl string
}

func (repo MongoDBSessionRepository) openDBConnection() mgo.Session {
	log.Println("connect to mongodb using url: " + repo.dbUrl)
	session, error := mgo.Dial(repo.dbUrl)
	if error != nil {
		panic(error)
	}
	session.SetMode(mgo.Monotonic, true)
	repo.db = session.DB(sessionDbName)
	return *session
}
