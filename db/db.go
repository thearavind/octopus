package db

import (
	"github.com/kapitol-app/octopus/config"
	"github.com/kapitol-app/octopus/logger"
	"gopkg.in/mgo.v2"
)


type Db struct {
	config config.Config
}

func (db *Db)getCollection(col DbCollection) {
	session, err := mgo.Dial(config.Config.Mongo.Url)
	if err != nil {
		logger.Log("Error: could not connect to mongo db at:", db.config.Mongo.Url)
		return nil
	}

	col := session.DB(db.config.Mongo.Db).C(db.config.Mongo.MembersCollection)

}

func (db *Db)Insert(T OctoModel) error {

}
