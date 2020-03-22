package repository

import "github.com/matgomes/stolen-bike-challenge/models"

func (r *Repository) FindOfficer() (models.Officer, error) {

	var result models.Officer

	err := db.C(OfficerCollection).Find(nil).One(&result)

	return result, err
}

func (r *Repository) InsertOfficer(o models.Officer) error {
	return db.C(OfficerCollection).Insert(o)
}

func (r *Repository) RemoveOfficer(officer models.Officer) error {
	return db.C(OfficerCollection).RemoveId(officer.Id)
}
