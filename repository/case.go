package repository

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/matgomes/stolen-bike-challenge/models"
)

func (r *Repository) GetAllCases() ([]models.Case, error) {

	var result []models.Case

	err := db.C(CaseCollection).Find(nil).All(&result)

	return result, err
}

func (r *Repository) GetCaseByID(id string) (models.Case, error) {

	var result models.Case

	isObjectIdHex := bson.IsObjectIdHex(id)
	if !isObjectIdHex {
		return result, errors.New("not id hex")
	}

	err := db.C(CaseCollection).FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

func (r *Repository) UpdateCase(c models.Case) error {
	return db.C(CaseCollection).UpdateId(c.Id, c)
}

func (r *Repository) InsertCase(c models.Case) error {
	c.Id = bson.NewObjectId()
	return db.C(CaseCollection).Insert(c)
}
