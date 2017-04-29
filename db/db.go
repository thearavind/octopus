package db

import (
	"github.com/greenac/octopus/config"
	"github.com/greenac/octopus/logger"
	"gopkg.in/mgo.v2"
)


type Db struct {
	config config.Config
}

func (db *Db)getSession() (session *mgo.Session) {
	session, err := mgo.Dial(config.Config.Mongo.Url)
	if err != nil {
		logger.Log("Error: could not connect to mongo db at:", db.config.Mongo.Url)
		return nil
	}

	return session

	col := session.DB(db.config.Mongo.Db).C(db.config.Mongo.MembersCollection)
}

func (db *Db)Insert(T OctoModel) error {

}
