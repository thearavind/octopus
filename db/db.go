package db

import (
	"github.com/greenac/octopus/config"
	"gopkg.in/mgo.v2"
)

type Db struct {
	config config.Config
}

func (db *Db)getSession() bool {
	session, err := mgo.Dial(config.Config.Mongo.Url)
	if err != nil {
		
	}
}
