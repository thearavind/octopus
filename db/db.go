package db

import (
	"errors"
	"github.com/kapitol-app/octopus/config"
	"github.com/kapitol-app/octopus/logger"
	"gopkg.in/mgo.v2"
)

type database struct {
	Session               *mgo.Session
	HouseCollection       *mgo.Collection
	SenatorCollection     *mgo.Collection
	LegislativeCollection *mgo.Collection
}

type CollectionType string

const (
	SenateCollection      CollectionType = "senate"
	HouseCollection       CollectionType = "house"
	LegislativeCollection CollectionType = "legislation"
)

//Connection - Global variable of the mongo db session struct
var db *database = nil

func initialize() {
	if db == nil {
		session, err := mgo.Dial(config.C.Mongo.Url)
		if err != nil {
			logger.Log("Error: could not connect to mongo db at:", config.C.Mongo.Url)
			panic("Error: could not connect to the database")
		}

		logger.Log("Initializing Mongo database")
		db = &database{
			Session:               session,
			HouseCollection:       session.DB(config.C.Mongo.Db).C(config.C.Mongo.HouseCollection),
			SenatorCollection:     session.DB(config.C.Mongo.Db).C(config.C.Mongo.SenatorsCollection),
			LegislativeCollection: session.DB(config.C.Mongo.Db).C(config.C.Mongo.LegislativeCollection),
		}
	}
}

func collectionForType(ct CollectionType) (*mgo.Collection, error) {
	var c *mgo.Collection
	switch ct {
	case SenateCollection:
		c = db.SenatorCollection
	case HouseCollection:
		c = db.HouseCollection
	case LegislativeCollection:
		c = db.LegislativeCollection
	default:
		return nil, errors.New("Collection Type Unknown")
	}

	return c, nil
}

func Find(m interface{}, col CollectionType) (*mgo.Query, error) {
	initialize()
	c, err := collectionForType(col)
	if err != nil {
		logger.Error("Failed to find query", m, "in collection:", col)
		return nil, err
	}

	q := c.Find(m)
	return q, nil
}

func Insert(m interface{}, col CollectionType) error {
	initialize()
	c, err := collectionForType(col)
	if err != nil {
		logger.Error("Failed to insert", m, "into collection:", col)
		return err
	}

	return c.Insert(m)
}

func Update(m interface{}, oId string, col CollectionType) error {
	initialize()
	c, err := collectionForType(col)
	if err != nil {
		logger.Error("Failed to update", m, "into collection:", col)
		return err
	}

	return c.UpdateId(oId, m)
}

func Upsert(m interface{}, q interface{}, col CollectionType) error {
	initialize()
	c, err := collectionForType(col)
	if err != nil {
		logger.Error("Failed to upsert model into collection:", col)
		return err
	}

	_, err = c.Upsert(q, m)
	return err
}
