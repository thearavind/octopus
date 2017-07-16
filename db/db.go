package db

import (
	"github.com/kapitol-app/octopus/config"
	"github.com/kapitol-app/octopus/logger"
	"gopkg.in/mgo.v2"
)

type connections struct {
	Session           *mgo.Session
	MemberCollection  *mgo.Collection
	SenatorCollection *mgo.Collection
}

//Connection - Global variable of the mongo db session struct
var Connection connections

func init() {
	session, err := mgo.Dial(config.C.Mongo.Url)
	if err != nil {
		logger.Log("Error: could not connect to mongo db at:", config.C.Mongo.Url)
		panic("Error: could not connect to the database")
	}

	Connection = connections{
		Session:           session,
		MemberCollection:  session.DB(config.C.Mongo.Db).C(config.C.Mongo.MembersCollection),
		SenatorCollection: session.DB(config.C.Mongo.Db).C(config.C.Mongo.SenatorsCollection),
	}
}
