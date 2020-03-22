package repository

import (
	"github.com/globalsign/mgo"
	"log"
)

const (
	CaseCollection    = "cases"
	OfficerCollection = "officers"
)

type Repository struct {
	Server   string
	Database string
}

var db *mgo.Database

func (r *Repository) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(r.Database)
}

func (r *Repository) Close() {
	db.Session.Close()
}
