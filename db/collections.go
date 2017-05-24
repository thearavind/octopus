package db

import "github.com/greenac/octopus/config"


type DbCollection string

type dbCollections struct {
	config *config.Config
}

func (col *dbCollections)Initialize() {
	if col.config == nil {
		c := config.Configuration()
		col.config = &c
	}
}

var dbCols dbCollections

func initializeDbCollections() {
	if dbCols == nil {
		dbCols = dbCollections{}
		dbCols.Initialize()
	}
}

func MembersCollection() DbCollection {
	initializeDbCollections()
	return dbCollections.config.Mongo.MembersCollection
}
